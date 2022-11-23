package lead

import(
	"github.com/jinzhu/gorm"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm/diaelcts/sqlite"
	"github.com/shashank/go-fiber-crm-basic/database"

)


type Lead struct{
	gorm.Model
	Name 		string		`json:"name"`
	Company 	string		`json:"company"`
	Email 		string		`json:"email"`
	phone 		int			`json:"phone"`
}

func GetLeads(c *fiber.Ctx){
	db := database.DBconn
	var leads []Lead
	db.Find(&lead, id)
	c.JSON(leads)
}


func GetLead( ){
	id := c.Params("id")
	db := database.DBconn
	var lead Lead 
	db.Find(&lead, id)
	c.JSON(lead)
}

func NewLead(){
	db := database.DBconn
	lead := new(lead)
	if err := c.BodyParser(lead); err !=nil {
		c.Status(503).Send(err)
		return 
	}
	db.Create(&lead)
	c.JSON(lead)
}

func DeleteLead(){
	id := c.Params("id")
	db := database.DBconn

	var lead Lead
	db.First(&lead, id)
	if lead,name == ""{
		c.Status(500),Send("no lead found with ID")
		return
	}
		
	}
	db.Delete(&lead)
	c.Send("lead successfully deleted")
}

