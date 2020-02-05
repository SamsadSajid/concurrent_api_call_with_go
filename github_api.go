package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

func getUserInfoFromGithub(userName string, wg *sync.WaitGroup) {
	fmt.Println("Executing getUserInfoFromGithub")

	resp, _ := http.Get("https://api.github.com/users/" + userName)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	githubUserInfo := GithubUserInfo{}
	json.Unmarshal(body, &githubUserInfo)

	fmt.Println(
		"Login Name: = ", githubUserInfo.Login,
		"Followers:=", githubUserInfo.Followers,
		//"Bio:=", githubUserInfo.Bio,
	)

	wg.Done()
}

func getFollowerListFromGithub(userName string, wg *sync.WaitGroup){
	fmt.Println("Executing getFollowerListFromGithub")

	resp, _ := http.Get("https://api.github.com/users/" + userName + "/followers")
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var githubUserFollowerList []Owner
	json.Unmarshal(body, &githubUserFollowerList)

	var w sync.WaitGroup

	fmt.Println(len(githubUserFollowerList))

	for i:=0; i < len(githubUserFollowerList); i++ {
		w.Add(1)
		fmt.Printf("%dth Follower's user name: %s\n", i, githubUserFollowerList[i].Login)

		go getUserInfoFromGithub(githubUserFollowerList[i].Login, &w)
	}

	w.Wait()

	wg.Done()
}

func getUserInfo(c *gin.Context) {
	fmt.Print(c.Param("userName"))
	startTime := time.Now()

	var wg sync.WaitGroup

	wg.Add(2)
	go getUserInfoFromGithub(c.Param("userName"), &wg)
	go getFollowerListFromGithub(c.Param("userName"), &wg)

	wg.Wait()

	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Hello",
	})
}

type GithubUserInfo struct {
	Login     string
	Followers int
}

type Owner struct {
	Login string
	Id string
	Url string
}