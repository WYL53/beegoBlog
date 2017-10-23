package models

import (
	"testing"
)

func TestTopic(t *testing.T) {
	if err := AddTopic("testTitle","cate","haha"); err != nil{
		t.Fatal(err)
	}
	topic,err := GetAllTopics(true)
	if len(topic) != 1 || err != nil || topic[0].Title != "testTitle"{
		t.Log("len(topic)=",len(topic))
		t.Log("err:",err)
		t.Fatal("GetAllCategories")
	}
	if err = ModifyTopic(topic[0].IdString(),"title_modify","cate2",topic[0].Content);err != nil{
		t.Log(err)
	}

	if err = DeleteTopic(topic[0].IdString());err != nil{
		t.Fatal(err)
	}
}


