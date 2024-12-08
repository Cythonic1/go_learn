package learnbyexample


import (
    "github.com/gin-gonic/gin"
    "net/http"
)


type Projects struct {
    ID int `json:"id"`;
    ProjectName string `json:"project_name"`;
    ProjectDesc string `json:"project_desc"`;
    ProjectURL string `json:"project_url"`;

}

func get_projects(c *gin.Context){
    var projects = [] Projects {
        {ID: 1, ProjectName: "Web scanner", ProjectDesc: "Rust based web testing tool", ProjectURL: "https://github.com/Cythonic1/web_tester"},
        {ID: 2, ProjectName: "Block chain", ProjectDesc: "A block chain from scrach using my most loved programming language", ProjectURL: "https://github.com/Cythonic1/a_blockchain_system"},
        {ID: 3, ProjectName: "Key Logger", ProjectDesc: "A simple client server based Key logger", ProjectURL: "https://github.com/Cythonic1/KeyLogger"},
    }
    c.IndentedJSON(http.StatusOK, projects);


}
func RestApiSimple(){

    router := gin.Default();
    router.GET("/Yasser", get_projects);
    router.Run(":8000");

}
