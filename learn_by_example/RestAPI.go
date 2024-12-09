package learnbyexample

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)


type Projects struct {
    ID int `json:"id"`;
    ProjectName string `json:"project_name"`;
    ProjectDesc string `json:"project_desc"`;
    ProjectURL string `json:"project_url"`;

}

var projects = [] Projects {
    {ID: 1, ProjectName: "Web scanner", ProjectDesc: "Rust based web testing tool", ProjectURL: "https://github.com/Cythonic1/web_tester"},
    {ID: 2, ProjectName: "Block chain", ProjectDesc: "A block chain from scrach using my most loved programming language", ProjectURL: "https://github.com/Cythonic1/a_blockchain_system"},
    {ID: 3, ProjectName: "Key Logger", ProjectDesc: "A simple client server based Key logger", ProjectURL: "https://github.com/Cythonic1/KeyLogger"},
}
func get_projects(c *gin.Context){
    c.IndentedJSON(http.StatusOK, projects);
}

func post_projects(c *gin.Context){
    var new_project Projects;
    fmt.Println(c);
    if err := c.ShouldBindJSON(&new_project); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()});
        return;
    }

   projects = append(projects, new_project); 
    c.JSON(http.StatusCreated, new_project);

}
func RestApiSimple(){

    router := gin.Default();
    router.GET("/Yasser", get_projects);
    router.POST("/Yasser", post_projects);
    router.Run(":8000");

}
