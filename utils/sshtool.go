/*
@Time : 2019/5/24 15:56
@Author : Tester
@File : 一条小咸鱼
@Software: GoLand
*/
package utils

import (
	"bytes"
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
)

type SSHModel struct {
	session     *ssh.Session
	writeCloser io.WriteCloser
	out         io.Reader
}

type SshBean struct {
	Client *ssh.Client
}

//获取ssh客户端实例
func GetSshBean(user, password, host string, port int16) (*SshBean, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		client       *ssh.Client
		err          error
	)

	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))
	hostKeyCallbk := func(hostname string, remote net.Addr, key ssh.PublicKey) error {
		return nil
	}
	clientConfig = &ssh.ClientConfig{
		User: user,
		Auth: auth,
		// Timeout:             30 * time.Second,
		HostKeyCallback: hostKeyCallbk,
	}
	// connet to ssh
	addr = fmt.Sprintf("%s:%d", host, port)

	if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}

	bean := &SshBean{client}

	return bean, nil
}

/**
获取一个ssh的客户端
*/
func (s *SshBean) GetClient() *ssh.Client {
	return s.Client
}

//获取session,始终获得最新的实例
func (s *SshBean) GetSession() (*ssh.Session, error) {
	client := s.Client
	session, err := client.NewSession()
	if err != nil {
		return nil, err
	}
	return session, nil
}

//func (s *SshBean) RunSsh(cmd string, id int, stdout *bytes.Buffer, stderr *bytes.Buffer) error {
//	client, err := s.GetClient(config.UserName, config.Password, config.HostName, config.Port)
//	if err != nil {
//		logs.Error(err)
//		return err
//	}
//	session, err := client.NewSession()
//	//stdinBuf, err := session.StdinPipe()
//	if err != nil {
//		logs.Error(err)
//	}
//	//cmdlist := strings.Split(cmd, ";")
//	defer session.Close()
//	session.Stdout = stdout
//	session.Stderr = stderr
//	timeout := make(chan bool, 1)
//	go func() {
//		time.Sleep(3 * time.Second) // sleep 3 second
//		timeout <- true
//	}()
//
//	ch := make(chan error, 1)
//	go func() {
//		ch <- session.Run(cmd)
//	}()
//	select {
//	case <-timeout:
//		fmt.Println("执行超时,强制返回!")
//	case <-ch:
//		fmt.Println("执行返回")
//	}
//
//	//session.Shell()
//	//f, err := session.CombinedOutput(cmd)
//	//err = session.Shell()
//	//err = session.Start(cmd)
//	if err != nil {
//		logs.Error(err)
//		return err
//	}
//	//fmt.Print(string(f))
//	//for _, c := range cmdlist {
//	//   c = c + "\n"
//	//   stdinBuf.Write([]byte(c))
//	//}
//	//err = session.Wait()
//	//logs.Error(stdout.String() + stderr.String())
//	//err = session.Run("echo 123456")
//	//ret, err := strconv.Atoi( str.Replace( stdOut.String(), "\n", "", -1 )  )
//	if err != nil {
//		return err
//	}
//
//	//	fmt.Printf("%d, %s\n", ret, stdErr.String() )
//	return nil
//}

// RFC 4254 Section 6.5.

//func (s *SshBean) SendShell(cmd string) {
//	req := execMsg{
//		Command: cmd,
//	}
//
//	//ok, err := s.ch.SendRequest("exec", true, Marshal(&req))
//	b, e := s.Session.SendRequest("exec", true, ssh.Marshal(&req))
//	if e != nil {
//		log.Println(e)
//	}
//	log.Print(b)
//}

/**
获取一个 封装后的 SSHSession 会话
单例模式,只获得一个session 会话
*/
//func GetSSHSession() (*SSHModel, error) {
//	mutex.Lock()
//	defer mutex.Unlock()
//	var e error
//	if sshModel == nil {
//		config := utils.GetSSHConfig()
//		if &config == nil {
//			log.Fatal("")
//		}
//		if session, closer, out, e := getSSHConnect(config.UserName, config.Password, config.HostName, config.Port); e == nil {
//			sshModel = &SSHModel{session, closer, out}
//			return sshModel, nil
//		}
//		return nil, e
//	}
//
//	return sshModel, nil
//}

func (model *SSHModel) Close() {
	err := model.writeCloser.Close()
	if err != nil {
		log.Fatal("关闭writeCloser流发生错误-->", err)
	}
	err = model.session.Close()
	if err != nil {
		log.Fatal("关闭sshSession发生错误-->", err)
	}
}

func (model *SSHModel) SendMessage(cmd string, stdOut, stdErr *bytes.Buffer) {

	//var stdOut1 []byte
	cmdList := strings.Split(cmd, ";")
	//session := model.sshService
	//defer model.Close()
	//session.Stdout = stdOut
	//session.Stderr = stdErr
	//closer:= model.closer
	for _, c := range cmdList {
		//c = c + "\n"
		//_, e := closer.Write([]byte(c))
		////output, e := session.CombinedOutput(c)
		//if e != nil{
		//	fmt.Println("写出ssh消息时发生错误-->",e)
		//}
		_ = model.session.Run(c)
	}
	//_, e := model.out.Read(stdOut1)
	//fmt.Println("-->",stdOut1)
	//fmt.Println("-->",e)
	//all, e := ioutil.ReadAll(model.out)
	//if e != nil{
	//	logs.Error(e)
	//}
	//fmt.Print(all)
}

//func GetSSHConnect(){
//
//}

func getSSHConnect(user, password, host string, port int16) (*ssh.Session, io.WriteCloser, io.Reader, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		client       *ssh.Client
		session      *ssh.Session
		err          error
		closer       io.WriteCloser
		out          io.Reader
	)
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))

	hostKeyCallbk := func(hostname string, remote net.Addr, key ssh.PublicKey) error {
		return nil
	}

	clientConfig = &ssh.ClientConfig{
		User: user,
		Auth: auth,
		// Timeout:             30 * time.Second,
		HostKeyCallback: hostKeyCallbk,
	}

	// connet to ssh
	addr = fmt.Sprintf("%s:%d", host, port)

	if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, nil, nil, err
	}
	// create session
	if session, err = client.NewSession(); err != nil {
		return nil, nil, nil, err
	}
	if closer, err = session.StdinPipe(); err != nil {
		return nil, nil, nil, err
	}
	if out, err = session.StdoutPipe(); err != nil {
		return nil, nil, nil, err
	}
	err = session.Shell()
	if err != nil {
		fmt.Println("打开shell发生错误-->", err)
	}
	return session, closer, out, nil
}

func (s *SshBean) GetSftpConnect() (sftpClient *sftp.Client, err error) { //参数: 远程服务器用户名, 密码, ip, 端口
	if sftpClient, err = sftp.NewClient(s.Client); err != nil { //创建客户端
		fmt.Println("创建客户端失败", err)
		return
	}

	return
}

/**
上传本地文件到远程服务器上
*/
func UploadFileRemote(sftpClient *sftp.Client, localFilePath string, remotePath string) {
	//打开本地文件流
	srcFile, err := os.Open(localFilePath)
	if err != nil {
		fmt.Println("os.Open error : ", localFilePath)
		log.Fatal(err)
	}
	//关闭文件流
	defer srcFile.Close()
	//上传到远端服务器的文件名,与本地路径末尾相同
	//var remoteFileName = path.Base(localFilePath)
	//打开远程文件,如果不存在就创建一个
	dstFile, err := sftpClient.Create(remotePath)

	if err != nil {
		log.Println("sftpClient.Create error : ", err, "远程地址->", remotePath)
		return
	}
	//关闭远程文件
	defer dstFile.Close()
	//读取本地文件,写入到远程文件中(这里没有分快穿,自己写的话可以改一下,防止内存溢出)
	ff, err := ioutil.ReadAll(srcFile)
	if err != nil {
		fmt.Println("ReadAll error : ", localFilePath)
		return

	}
	dstFile.Write(ff)
}
