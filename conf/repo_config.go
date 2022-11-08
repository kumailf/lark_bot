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

func UpdateRepoConf(repo string, key string, value string) {
	flag := utils.MGDBUpdateOne(db, repo, "repo", repo, key, value)
	if flag {
		RepoConf[repo] = utils.MGDBFindOne(db, repo, "repo", repo)
	}
}

func FuncIsWork(repo string, func_name string) bool {
	repo_conf, ok := RepoConf[repo]
	if !ok {
		return true
	}
	tmp, ok := repo_conf[func_name]
	if !ok {
		return true
	}
	if tmp.(string) == "off" {
		return false
	}
	return true
}
