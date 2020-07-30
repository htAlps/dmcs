//  ._______.___________________.___________________.___________________.___________________.___________________.___________________.___________________.                   .                   .                   .                   .                   ;
//  mux.go  ▻20‹05‹07τ10›25›21›     
    package lib

    import (
        "fmt"; "log"; "time"; "context"; "encoding/json"; "net/http"; "github.com/gorilla/mux";
        prmtv "go.mongodb.org/mongo-driver/bson/primitive"; "go.mongodb.org/mongo-driver/mongo";
        "go.mongodb.org/mongo-driver/mongo/options";
        "go.mongodb.org/mongo-driver/bson";
    )

    // CONSTANTS: CRUD                      RESTful METHODS                         MONGO CRUD
    const create    string  = `create`;     const M_post    string  = `POST`;       const insert    string  = `insert`;
    const read      string  = `read`;       const M_get     string  = `GET`;        const find      string  = `find`;
    const update    string  = `update`;     const M_put     string  = `PUT`;        // const update string  = `update`;
    const delete    string  = `delete`;     const M_del     string  = `DELETE`;     const remove    string  = `delete`;

    // SCHEMAS / DBs
    const db_name           string  = `mongodb`                         //  schema must be "mongodb" or "mongodb+srv" 
    const db_url            string  = db_name + `://localhost:27017`

    // Entities
    const e1_person         string  = `person`

    // Paths
    const p1_postPerson     string  = `/` + e1_person
    const p2_putPerson      string  = `/` + e1_person + `/{id}`
    const p3_getPersons     string  = `/` + e1_person + `s`
    const p4_getPerson      string  = `/` + e1_person + `/{id}`
    const p5_delPerson      string  = `/` + e1_person + `/{id}`

    // REST Ports
    const r1_port           string  = `:8051`


//  ._______.___________________.___________________.___________________.___________________.___________________._______;
//  Data Model - Types

    type Person struct {
        ID      prmtv.ObjectID  `json:"_id,omitempty"   bson:"_id,omitempty"`
        Name    string          `json:"name,omitempty"  bson:"name,omitempty"`
        Email   string          `json:"email,omitempty" bson:"email,omitempty"`
    }

    var cliDB *mongo.Client

    // Handle1_postPerson creates a new record in DB Collection: Person
    func H1_postPerson(resp http.ResponseWriter, req *http.Request) {
        fmt.Printf("\nStarting H1_postPerson\n\n")
        resp.Header().Set("content_type", "application/json")

        var record Person
        _        = json.NewDecoder(req.Body).Decode(&record)
        fmt.Printf("\nRecord: %v\n", record)

        coll     := cliDB.Database(db_name).Collection(e1_person)
        ctx, cancel    := context.WithTimeout(context.Background(), 10 * time.Second)
        if cancel != nil {
            fmt.Printf("\nCONTEXT ERROR: %v \n", cancel)
        }

        res, _  := coll.InsertOne(ctx, record)

        json.NewEncoder(resp).Encode(res)
    }


    // Handle2_putPerson Updates a record in DB Collection: Person using ID
    func H2_putPerson(resp http.ResponseWriter, req *http.Request) {
        fmt.Printf("\nStarting H2_putPerson\n\n")

    }


    // Handle3_getPersons selects all records in DB Collection: Person
    func H3_getPersons(resp http.ResponseWriter, req *http.Request) {
        fmt.Printf("\nStarting H3_getPersons\n\n")
        resp.Header().Set("content_type", "application/json")

        var records []Person

        coll     := cliDB.Database(db_name).Collection(e1_person)
        ctx, cancel    := context.WithTimeout(context.Background(), 10 * time.Second)
        if cancel != nil {
            fmt.Printf("\nCONTEXT ERROR: %v \n", cancel)
        }

        cursor, _   := coll.Find(ctx, bson.M{})

        defer cursor.Close(ctx)

        for cursor.Next(ctx) {
            var record Person
            cursor.Decode(&record)
            records = append(records, record)
            fmt.Printf("\nRecord: %v\n", record)
        }
        json.NewEncoder(resp).Encode(records)
    }


    // Handle4_getPerson finds a record in DB Collection: Person using ID
    func H4_getPerson(resp http.ResponseWriter, req *http.Request) {
        fmt.Printf("\nStarting H4_getPerson\n\n")

    }


    // Handle5_delPerson deletes a record in DB Collection: Person
    func H5_delPerson(resp http.ResponseWriter, req *http.Request) {
        fmt.Printf("\nStarting H5_delPerson\n\n")

    }


//  ._______.___________________.___________________.___________________.___________________.___________________._______;
//  http Clients / services 
    // SvcA_mongo_rest_dmcs_intest means: web service#A, Mongo DB, REST API, itest=Integration Test
    // this function integration tests a basic CRUD API on Mongo 

    func ITestA_rest_dmcs(key string) {
        fmt.Printf("Key: %s \n", key)
    }

//  ._______.___________________.___________________.___________________.___________________.___________________._______;
//  http servers / services 

    func SvcA_mongo_rest_dmcs() {

        fmt.Printf("Starting \n")
        ctx, cancel    := context.WithTimeout(context.Background(), 10 * time.Second)
        if cancel != nil {
            fmt.Printf("\nCONTEXT ERROR: %v \n", cancel)
        }
        cliOptions  := options.Client().ApplyURI(db_url)
        var err error
        cliDB, err  = mongo.Connect(ctx, cliOptions)
        if err != nil {
            fmt.Printf("\nERROR: %v \n", err)
            log.Fatal(err)
        }
        rr := mux.NewRouter()

        rr.HandleFunc("/person",      H1_postPerson).Methods(M_post)
        rr.HandleFunc("/person/{id}", H2_putPerson).Methods(M_put)
        rr.HandleFunc("/persons",     H3_getPersons).Methods(M_get)
        rr.HandleFunc("/person/{id}", H4_getPerson).Methods(M_get)
        rr.HandleFunc("/person/{id}", H5_delPerson).Methods(M_del)

/*
        rr.HandleFunc(p1_postPerson, H1_postPerson).Methods(M_post)
        rr.HandleFunc(p2_putPerson,  H2_putPerson).Methods(M_put)
        rr.HandleFunc(p3_getPersons, H3_getPersons).Methods(M_get)
        rr.HandleFunc(p4_getPerson,  H4_getPerson).Methods(M_get)
        rr.HandleFunc(p5_delPerson,  H5_delPerson).Methods(M_del)
*/
        fmt.Printf("p1_postPerson: %s \np2_putPerson: %s \np3_getPersons: %s \np4_getPerson: %s \np5_delPerson: %s \n",
                    p1_postPerson,      p2_putPerson,      p3_getPersons,      p4_getPerson,      p5_delPerson)
        fmt.Printf("M_post: %s \nM_put: %s \nM_get: %s \n", M_post, M_put, M_get)

        http.ListenAndServe(":8051", rr)
        // http.ListenAndServe(r1_port, rr)
    }



