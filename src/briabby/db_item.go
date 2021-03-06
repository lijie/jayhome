package briabby

import (
	"errors"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"os"
)

// ProtoItem indicate an item that saved in db
type ProtoItem struct {
	ID         string   `bson:"_id"`
	Name       string   `bson:"name"`
	Promotion  int      `bson:"promotion"`
	ImageSmall string   `bson:"smallimage"`
	ImageBig   string   `bson:"bigimage"`
	Desc       string   `bson:"desc"`
	Price      []string `bson:"price"`
	PaypalBtn  string   `bson:"paypalbtn"`
	Category   string   `bson:"category"`
}

type Store struct {
	session   *mgo.Session
	db        *mgo.Database
	itemTable *mgo.Collection
}

func (s *Store) FindItem(id string) (*ProtoItem, error) {
	var result ProtoItem
	if err := s.itemTable.Find(bson.M{"_id": id}).One(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *Store) SaveItem(item *ProtoItem) error {
	if len(item.ID) == 0 {
		item.ID = bson.NewObjectId().Hex()
	}
	_, err := s.itemTable.Upsert(bson.M{"_id": item.ID}, item)
	return err
}

func (s *Store) DelItem(id string) error {
	var (
		p1   string
		p2   string
		err  error
		item *ProtoItem
	)
	if item, err = s.FindItem(id); err != nil {
		return err
	} else {
		p1 = item.ImageSmall
		p2 = item.ImageBig
	}
	if err = s.itemTable.Remove(bson.M{"_id": id}); err != nil {
		return err
	}
	os.Remove(".." + p1)
	os.Remove(".." + p2)
	return nil
}

func (s *Store) FindItemByCat(cat string) ([]ProtoItem, error) {
	var results []ProtoItem
	err := s.itemTable.Find(bson.M{"category": cat}).Sort("-promotion", "-_id").All(&results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (s *Store) All() ([]ProtoItem, error) {
	var results []ProtoItem
	err := s.itemTable.Find(nil).All(&results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func NewStore(url string) (*Store, error) {
	var err error
	s := &Store{}
	if s.session, err = mgo.Dial(url); err != nil {
		return nil, err
	}
	if s.db = s.session.DB("babegarden"); s.db == nil {
		s.session.Close()
		return nil, errors.New("database not found")
	}
	if s.itemTable = s.db.C("item"); s.itemTable == nil {
		s.session.Close()
		return nil, errors.New("table not found")
	}
	return s, nil
}
