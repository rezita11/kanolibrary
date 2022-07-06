package routes


	import "encoding/json"
	import "go.mongodb.org/mongo-driver/bson"
	import "kanolibrary/crud"
	import "kanolibrary/models"
	import "kanolibrary/util"
	import "net/http"
	import "go.mongodb.org/mongo-driver/bson/primitive"

//get all
func FindAllBooks(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Conte-Type", "apllication/json")
	if r.Method == "GET" {
		result, err := crud.FindAll("books")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return
	}
http.Error(w, "Method not allowed", http.StatusInternalServerError)
}


//get by id
func FindByIdBook(w http.ResponseWriter, r *http.Request) {
w.Header().Set("Content-Type", "application/json")

		if r.Method == "GET" {
		requestId := r.URL.Query().Get("id")

		objId, err := primitive.ObjectIDFromHex(requestId)
		result, err := crud.FindOne("books", bson.M{"_id": objId})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		return

		}
		w.Write(result)
		return
	}
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)

}
//post
func InsertBook(w http.ResponseWriter, r *http.Request) {
w.Header().Set("Content-Type", "application/json")
 
	if r.Method == "POST"{
		var book models.Books
		err := json.NewDecoder(r.Body).Decode(&book)
		if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
		return
		}
		result, err := 
		crud.Insert("books", book)
		if err != nil {}
		w.Write([]byte(result))
		return
	} 
	http.Error(w ,"Method not allowed", http.StatusInternalServerError)
}

//update 
func UpdateBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Conte n-Type", "application/json")
	if r.Method == "POST"{
		var book models.Books
		err := json.NewDecoder(r.Body).Decode(&book)
		if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
		return
		}
		requestId := r.URL.Query().Get("id")

		objId, err := primitive.ObjectIDFromHex(requestId)
		util.ErrorChecker(err)     
		result, err := crud.Update("books", bson.M{"_id": objId}, bson.M{"$set":bson.M{"name":book.Name, "author":book.Author, "publisher":book.Publisher, "page":book.Page}})

		if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
		return
		}
		w.Write([]byte(result))
		return

	} 
	
http.Error(w ,"Method not allowed", http.StatusInternalServerError)
}
//deleete
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

		if r.Method == "DELETE" {
			requestId := r.URL.Query().Get("id")

		objId, err := primitive.ObjectIDFromHex(requestId)
		util.ErrorChecker(err)
		result, err := crud.Remove("book", bson.M{"_id": objId})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return

		}
		w.Write([]byte(result))
		return


	http.Error(w, "Method not allowed", http.StatusInternalServerError)
}
}