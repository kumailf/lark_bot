package conf

import "code.byted.org/larkim/oapi_demo/utils"

var RepoConf map[string]map[string]interface{}

const db = "lark_bot"

// how to use
// import conf
// func_conf := conf[repo][func].(string)
// output: on/off
func init() {
	RepoList := [...]string{"MyPic", "mmediting"}
	RepoConf = make(map[string]map[string]interface{})

	for _, repo := range RepoList {
		tmp := utils.MGDBFindOne(db, repo, "repo", repo)
		RepoConf[repo] = tmp
	}
}

func UpdateRepoConf(repo string, fcun string) {

}
