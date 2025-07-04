package main

import (
    "io"
	"math/rand"
    "log"
    "net/http"
    "time"
	"reflect"
	"strconv"
	"strings"
	"regexp"
"encoding/json"
    "fmt"
    "io/ioutil"
    "os"
	"unicode/utf8"
	
)

const (
    htmlTagStart = 60 // Unicode `<`
    htmlTagEnd   = 62 // Unicode `>`
	logger = true
	
)

// Aggressively strips HTML tags from a string.
// It will only keep anything between `>` and `<`.
func stripHtmlTags(s string) string {
    // Setup a string builder and allocate enough memory for the new string.
    var builder strings.Builder
    builder.Grow(len(s) + utf8.UTFMax)

    in := false // True if we are inside an HTML tag.
    start := 0  // The index of the previous start tag character `<`
    end := 0    // The index of the previous end tag character `>`

    for i, c := range s {
        // If this is the last character and we are not in an HTML tag, save it.
        if (i+1) == len(s) && end >= start {
            builder.WriteString(s[end:])
        }

        // Keep going if the character is not `<` or `>`
        if c != htmlTagStart && c != htmlTagEnd {
            continue
        }

        if c == htmlTagStart {
            // Only update the start if we are not in a tag.
            // This make sure we strip out `<<br>` not just `<br>`
            if !in {
                start = i
            }
            in = true

            // Write the valid string between the close and start of the two tags.
            builder.WriteString(s[end:start])
            continue
        }
        // else c == htmlTagEnd
        in = false
        end = i + 1
    }
    s = builder.String()
    return s
}

type User struct {
	Quiz   string
	Index  int `json:",omitempty"`
	Active bool
	Return string
}

// Users struct which contains
// an array of users
type Intents struct {
	Intents []chat `json:"intents"`
}

type chat struct {
	Tag       string
	Patterns  []string
	Responses []string
	Poid       string
	Context   []string
}



//Define a map to implement routing table.
var mux map[string]func(http.ResponseWriter , *http.Request) 

func main(){
	


// Creation Serveur WEB	
    server := http.Server{
        Addr: ":8091",
        Handler: &myHandler{},
        ReadTimeout: 5*time.Second,
    }
    
    mux = make(map[string]func(http.ResponseWriter, *http.Request))
    mux["/tmp"] = Tmp
    err := server.ListenAndServe()
    if err != nil {
        log.Fatal(err)
    }   
}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){

strArray := [6]string{"India", "Canada", "Japan", "Germany", "Italy", "Coucou"}
nameLength := len(strArray)

// Creation du JSon
userJson := `[{"Quiz": "le", "Index": 1, "Active": true, "Return": "un"},{"Quiz": "Coucou", "Index": 2, "Active": true, "Return": "Salut"},{"Quiz": "Hihi", "Index": 3, "Active": true, "Return": "Hi!"},{"Quiz": "nouvelle", "Index": 4, "Active": true, "Return": "news"},{"Quiz": "Cool", "Index": 5, "Active": true, "Return": "Tranquille"}]`
var users []User
// Lecture de JSon
json.Unmarshal([]byte(userJson), &users)
nameLength2 := len(users)


//Load Json file*****************************************************
// Open our jsonFile
jsonFile, err := os.Open("chatbot.json")
// if we os.Open returns an error then handle it
if err != nil {
    fmt.Println(err)
}
if (logger){
	fmt.Println("Successfully Opened chatbot.json")
}
// defer the closing of our jsonFile so that we can parse it later on
defer jsonFile.Close()
byteValue, _ := ioutil.ReadAll(jsonFile)
// we initialize our Users array
var intents Intents
// we unmarshal our byteArray which contains our
// jsonFile's content into 'intents' which we defined above
json.Unmarshal(byteValue, &intents)

    //fmt.Println(result["intents"])
nameLength3 := len(intents.Intents)


//**********************************************************************	

    // Implement route forwarding
    if h, ok := mux[r.URL.String()];ok{
    //Implement route forwarding with this handler, the corresponding route calls the corresponding func.
        h(w, r)
        return
    }
	// On supprimer la barre
	res0 := strings.Trim(r.URL.String(), "/")
	// On supprimer les espace HTML
	res1 := strings.Trim(res0, "%")
	// On supprime les espaces
	//whitespaces := regexp.MustCompile(`\s+`)
	//res2 := whitespaces.ReplaceAllString(res1, " ")
	// On supprime les chiffres "[0-9]"
	number := regexp.MustCompile(`[0-9]`)
	res10 := number.ReplaceAllString(res1, " ")
	
	rechercher := stripHtmlTags(res10)
	// On decoupe le text en plusieurs mots
	zp := regexp.MustCompile(` *, *`)             // spaces and one comma - espaces et une virgule
	rs := zp.Split(rechercher, -1) // ["a" "b" "c "]
	
	//init page:
	io.WriteString(w, "<!DOCTYPE html><html lang=\"fr\"><head><meta charset=\"UTF-8\"><meta http-equiv=\"X-UA-Compatible\" content=\"IE=edge\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><title>Chat Neoray</title></head><body>")
	io.WriteString(w, "<div style=\"white-space:pre-wrap;\">")
	
	// init:
		io.WriteString(w, "URL: "+rechercher)
		//io.WriteString(w, "\n"+"Nombre de mots actuellement connus: "+strconv.Itoa(nameLength)+"\n");
		//io.WriteString(w, "\n"+"Nombre de mots actuellement connus: "+strconv.Itoa(nameLength2)+"\n");
		io.WriteString(w, "\n"+"Nombre de mots actuellement connus: "+strconv.Itoa(nameLength3)+"\n");
	if (logger){
		fmt.Println("URL: " + rechercher)
		fmt.Println("\n" + "Nombre de mots actuellement connus: " + strconv.Itoa(nameLength))
		fmt.Println("\n" + "Nombre de mots actuellement connus: " + strconv.Itoa(nameLength2))
		fmt.Println("\n" + "Nombre de mots actuellement connus: " + strconv.Itoa(nameLength3))
	}	
	
	var loopq bool
	loopq = false
	
	for q := range rs {
		io.WriteString(w, "<!-- reqs: "+rs[q]+" -->")
		// On recherche l'ensemble des resultats si loopq est true
		resultat := QuizJson(intents, rs[q], loopq)
		fmt.Println("reqs: " + rs[q])
		//Si le resultat est eguale a ****, on ne traite pas la reponse, sinon :
		if (strings.Contains(resultat[0],"***") == false){	
			// Si la reopnse est boucle on considère que le prochain argument on attend une boucle.
			if strings.EqualFold(resultat[0], "*loops*"){
				io.WriteString(w, "<!-- *loops* -->")
				loopq = true
			}else{
				fmt.Println("\n" + "chatBOT: " + resultat[0])
			    io.WriteString(w, "\n"+""+ resultat[0]) //reponse chatBOT:
				loopq = false
			}	
			 
		}
		// Si le poid de la reponse est de 2 on consider que c'est une reponse final
		if (strings.Contains(resultat[1],"2") == true){ 
		break // break here
		}
	}
	
	//fmt.Println("\n" + "reponse final chatBOT: " + finalreponse)
	//fmt.Println(strings.Replace("le chemin de l'ecole", "le ", "un ", -1)) //un chemin de l'ecole
	io.WriteString(w, "</div>")
	io.WriteString(w, "</body></html>")
}
func printQuiz(users []User) {
	// Loop over structs and display them
	for l := range users {
		fmt.Printf("Question = %v, reponse = %v", users[l].Quiz, users[l].Return)
		fmt.Println()
	}
}

func QuizJson(intents Intents, item interface{}, loopq bool) []string {
	
	
	//breakpoid := "0"
	valStr := fmt.Sprint(item)
	for i := 0; i < len(intents.Intents); i++ {

		if strings.EqualFold(intents.Intents[i].Context[0],"fr"){
			if strings.Contains(intents.Intents[i].Tag, valStr) {
				fmt.Println("tag : " + intents.Intents[i].Tag)
				//Question clee:
				if (loopq) {
					fmt.Println("Question clee 1 ")
					return  []string{loopquiz (intents.Intents[i].Responses), intents.Intents[i].Poid}
				}else{	
					fmt.Println("Context : " + intents.Intents[i].Context[0])
					fmt.Println("Poid : " + intents.Intents[i].Poid)
					vv := randint(len(intents.Intents[i].Responses))
					return  []string{intents.Intents[i].Responses[vv], intents.Intents[i].Poid}
				}
			}
		}
	}

	for i := 0; i < len(intents.Intents); i++ {
		if strings.EqualFold(intents.Intents[i].Context[0],"fr"){
			for u := 0; u < len(intents.Intents[i].Patterns); u++ {
				if strings.Contains(intents.Intents[i].Patterns[u], valStr) {
									//Question clee:
				if (loopq) {
					fmt.Println("Question clee 1 ")
					return []string{loopquiz (intents.Intents[i].Responses), intents.Intents[i].Poid}
				}else{	
					if (logger){
						fmt.Println("tag : " + intents.Intents[i].Tag)
						fmt.Println("Context : " + intents.Intents[i].Context[0])
						fmt.Println("Poid : " + intents.Intents[i].Poid)
					}
					vv := randint(len(intents.Intents[i].Responses))
					return []string{intents.Intents[i].Responses[vv], intents.Intents[i].Poid}
				}
				}
			}
		}
	}
	return []string{"", "0"}
}

// On donne toutes les reponses:
func loopquiz(reponses []string) string {
	var str1 string
	str1 = ""

	for i := 0; i < len(reponses); i++ {
		str1 = str1+ "\n"+strconv.Itoa(i)+" - " + reponses[i]
	}	
	return str1
}	

func Quiz(users []User, item interface{}) string {

	for l := range users {
		if users[l].Quiz == item {
			return users[l].Return
		}
	}
	return "non"
}

func Tmp(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "version 3")
}

func itemExists(arrayType interface{}, item interface{}) string {
	arr := reflect.ValueOf(arrayType)

	if arr.Kind() != reflect.Array {
		panic("Invalid data-type")
	}

	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == item {
			return "oui"
		}
	}

	return "non"
}

func randint(vmax int) int {

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return r1.Intn(vmax)
}

