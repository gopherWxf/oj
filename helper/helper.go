package helper

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"net/smtp"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"oj/define"

	"github.com/dgrijalva/jwt-go"
	"github.com/jordan-wright/email"
	uuid "github.com/satori/go.uuid"
)

type UserClaims struct {
	Identity string `json:"identity"`
	Name     string `json:"name"`
	IsAdmin  int    `json:"is_admin"`
	jwt.StandardClaims
}

// GetMd5
// 生成 md5
func GetMd5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

var myKey = []byte("gin-gorm-oj-key")

// GenerateToken
// 生成 token
func GenerateToken(identity, name string, isAdmin int) (string, error) {
	UserClaim := &UserClaims{
		Identity:       identity,
		Name:           name,
		IsAdmin:        isAdmin,
		StandardClaims: jwt.StandardClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim)
	tokenString, err := token.SignedString(myKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// AnalyseToken
// 解析 token
func AnalyseToken(tokenString string) (*UserClaims, error) {
	userClaim := new(UserClaims)
	claims, err := jwt.ParseWithClaims(tokenString, userClaim, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return nil, fmt.Errorf("analyse Token Error:%v", err)
	}
	return userClaim, nil
}

// SendCode
// 发送验证码
func SendCode(toUserEmail, code string) error {
	e := email.NewEmail()
	e.From = "Wxf <68725032@qq.com>"
	e.To = []string{toUserEmail}
	e.Subject = "验证码发送"
	e.HTML = []byte("您的验证码：<b>" + code + "</b>")
	// 返回 EOF 时，关闭SSL重试
	return e.Send("smtp.qq.com:587",
		smtp.PlainAuth("", "68725032@qq.com", define.MailPassword, "smtp.qq.com"),
	)
}

// GetUUID
// 生成唯一码
func GetUUID() string {
	return uuid.NewV4().String()
}

// GetRand
// 生成验证码
func GetRand() string {
	rand.NewSource(time.Now().UnixNano())
	s := ""
	for i := 0; i < 6; i++ {
		s += strconv.Itoa(rand.Intn(10))
	}
	return s
}

// CodeSave
// 保存代码
func CodeSave(code []byte) (string, error) {
	dirName := "code/" + GetUUID()
	path := dirName + "/main.go"
	err := os.Mkdir(dirName, 0777)
	if err != nil {
		return "", err
	}
	f, err := os.Create(path)
	if err != nil {
		return "", err
	}
	f.Write(code)
	defer f.Close()
	return path, nil
}

// CheckGoCodeValid
// 检查golang代码的合法性
func CheckGoCodeValid(path string) (bool, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return false, err
	}
	code := string(b)
	for i := 0; i < len(code)-6; i++ {
		if code[i:i+6] == "import" {
			var flag byte
			for i = i + 7; i < len(code); i++ {
				if code[i] == ' ' {
					continue
				}
				flag = code[i]
				break
			}
			if flag == '(' {
				for i = i + 1; i < len(code); i++ {
					if code[i] == ')' {
						break
					}
					if code[i] == '"' {
						t := ""
						for i = i + 1; i < len(code); i++ {
							if code[i] == '"' {
								break
							}
							t += string(code[i])
						}
						if _, ok := define.ValidGolangPackageMap[t]; !ok {
							return false, nil
						}
					}
				}
			} else if flag == '"' {
				t := ""
				for i = i + 1; i < len(code); i++ {
					if code[i] == '"' {
						break
					}
					t += string(code[i])
				}
				if _, ok := define.ValidGolangPackageMap[t]; !ok {
					return false, nil
				}
			}
		}
	}
	return true, nil
}

func ToTime(num int64) time.Time {
	return time.Unix(num, 0)
}
func GetMemory(pid string) int {
	// 执行ps命令获取进程内存使用情况
	output, err := exec.Command("ps", "-o", "rss=", "-p", pid).Output()
	if err != nil {
		return -1
	}
	// 解析输出结果，并转换为整数（单位为KB）
	memoryUsage := strings.TrimSpace(string(output))
	if usage, err := strconv.Atoi(memoryUsage); err == nil {
		return usage / 1024
	} else {
		return -1
	}
}

func KillProcess(pid string) {
	cmd := exec.Command("kill", "-9", pid)
	cmd.Run()
}
func Max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
