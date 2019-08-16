package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"time"

	pb "github.com/mclellac/ramify/services/post"

	"github.com/go-yaml/yaml"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var conf Config

type Config struct {
	Domain        string `yaml:"domain"`
	Port          string `yaml:"port"`
	ServerLogfile string `yaml:"server_logfile"`
	PubKey        string `yaml:"pub_key"`
	CertPEM       string `yaml:"cert_pem"`
	CertCSR       string `yaml:"cert_csr"`
	CertCRT       string `yaml:"cert_crt"`
	TypeDB        string `yaml:"type_db"`
	DBConn        string `yaml:"db_connect"`
	DBUsername    string `yaml:"db_username"`
	DBPassword    string `yaml:"db_password"`
	DBHostname    string `yaml:"db_hostname"`
	DBName        string `yaml:"db_name"`
	DBLogfile     string `yaml:"db_logfile"`
}

type Post struct {
	ID      int64  `gorm:"primary_key"`
	Created int32  `gorm:"size:25"`
	Title   string `gorm:"type:varchar(100)"`
	Article string `gorm:"type:varchar(5000)"`
}

type postService struct {
	post []*pb.Content
	m    sync.Mutex
	DB   *gorm.DB
}

func (ps *postService) Delete(ctx context.Context, req *pb.Content) (*pb.Response, error) {
	ps.m.Lock()
	defer ps.m.Unlock()

	if ps.DB.First(req).RecordNotFound() {
		return &pb.Response{
			Error: fmt.Sprintf("Sorry, chummer. I can't find a post with %d as an ID.", int64(req.Id)),
		}, nil
	} else {
		ps.DB.Delete(req)
		return &pb.Response{
			Message: fmt.Sprintf("Removed post with ID %d.", int64(req.Id)),
		}, nil
	}
}

func (ps *postService) Add(ctx context.Context, req *pb.Content) (*pb.Response, error) {
	ps.m.Lock()
	defer ps.m.Unlock()
	ps.post = append(ps.post, req)

	req.Created = int32(time.Now().Unix())
	ps.DB.Save(&req)

	return &pb.Response{
		Message: fmt.Sprintf("Post with title '%s' submitted.", req.Title),
	}, nil
}

func (ps *postService) List(req *pb.Request, stream pb.Post_ListServer) error {
	var post []*pb.Content

	ps.m.Lock()
	defer ps.m.Unlock()

	ps.DB.Order("created desc").Find(&post)

	for _, r := range post {
		if err := stream.Send(r); err != nil {
			return err
		}
	}
	return nil
}

func (c *Config) Init() {
	file, err := os.Open("postd.yaml")
	if err != nil {
		log.Fatalf("failed to open postd.yaml server config file: %v", err)
	}
	defer file.Close()

	stat, _ := file.Stat()

	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		log.Fatalf("failed to open settings file: %v", err)
	}

	err = yaml.Unmarshal(bs, &c)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}

func main() {
	conf.Init()

	connectionString := "postgres://" + conf.DBUsername + ":" +
		conf.DBPassword + "@" +
		conf.DBHostname + ":5432/" +
		conf.DBName + "?sslmode=disable"

	db, err := gorm.Open(conf.TypeDB, connectionString)
	if err != nil {
		log.Fatalf("failed to open DB: %v", err)
	}
	defer db.Close()

	dblog, err := os.OpenFile(conf.DBLogfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer dblog.Close()

	// Disable table name pluralization
	db.SingularTable(true)
	// Enable Logger
	db.LogMode(true)
	db.SetLogger(log.New(dblog, "\r\n", 0))

	if !db.HasTable("post") {
		//db.AutoMigrate(&Post{})
		db.CreateTable(&Post{})
	}

	lis, err := net.Listen("tcp", conf.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterPostServer(s, &postService{DB: db})
	fmt.Println("Server started on port", conf.Port)
	s.Serve(lis)
}
