package main
 
import (
   "context"
   "log"
 
   "github.com/Pichais/app/controllers"
   _ "github.com/Pichais/app/docs"
   "github.com/Pichais/app/ent"
   "github.com/gin-contrib/cors"
   "github.com/gin-gonic/gin"
   _ "github.com/mattn/go-sqlite3"
   swaggerFiles "github.com/swaggo/files"
   ginSwagger "github.com/swaggo/gin-swagger"
)

type Organs struct {
	Organ []Organ
}
type Organ struct {
	OrganName  string
}


type TypeDiseases struct {
	TypeDisease []TypeDisease
}
type TypeDisease struct {
	DiseaseName  string	
}

type Physicians struct {
	Physician []Physician
}
type Physician struct {
    PhysicianName string
    PhysicianEmail string
}

// @title SUT SA Example API
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/
 
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
 
// @host localhost:8080
// @BasePath /api/v1
 
// @securityDefinitions.basic BasicAuth
 
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
 
// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information
 
// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information
 
// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information
 
// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information

func main() {
   router := gin.Default()
   router.Use(cors.Default())
 
   client, err := ent.Open("sqlite3", "file:ent.db?&cache=shared&_fk=1")
   if err != nil {
       log.Fatalf("fail to open sqlite3: %v", err)
   }
   defer client.Close()
 
   if err := client.Schema.Create(context.Background()); err != nil {
       log.Fatalf("failed creating schema resources: %v", err)
   }
 
   v1 := router.Group("/api/v1")
    controllers.NewOrganController(v1, client)
    controllers.NewPhysicianController(v1, client)
    controllers.NewSpacialityController(v1, client)
    controllers.NewTypeDiseaseController(v1, client)
  
    organs := Organs{
		Organ: []Organ{
			Organ{"eye (ตา)"},
            Organ{"Heart (หัวใจ)"},
            Organ{"Bone (กระดูก)"},
		},
	}

	for _, o := range organs.Organ {
		client.Organ.
			Create().
			SetOrganName(o.OrganName).
			Save(context.Background())
	}

    // Set Users Data
	typeDiseases := TypeDiseases{
		TypeDisease: []TypeDisease{
			TypeDisease{"Cataract (ต้อกระจก)"},
			TypeDisease{"Glaucoma (ต้อหิน)"},
			TypeDisease{"Heart arrhythmia (หัวใจเต้นผิดจังหวะ)"},
			TypeDisease{"Vulvular Heart Disease (โรคลิ้นหัวใจ)"},
			TypeDisease{"Abnormal joints (ข้อต่อผิดปกติ)"},
			TypeDisease{"หมอนรองกระดูกสันหลัง กดทับเส้นประสาท"},
		},
	}

	for _, td := range typeDiseases.TypeDisease {
		client.TypeDisease.
			Create().
			SetDiseaseName(td.DiseaseName).
			Save(context.Background())
    }
    
    physicians := Physicians{
		Physician: []Physician{
			Physician{"Kanjana sangtong", "Kanjana@love.co.th"},
			Physician{"Sittisak kombai", "Sittisak@jipjip.com"},
			Physician{"Songpol sangtong", "Songpol@love.co.th "},
		},
	}

	for _, p := range physicians.Physician {
		client.Physician.
			Create().
			SetPhysicianName(p.PhysicianName).
			SetPhysicianEmail(p.PhysicianEmail).
			Save(context.Background())
	}

   router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
   router.Run()
}
