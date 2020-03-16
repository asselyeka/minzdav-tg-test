package main
import (
  "net/http"
  "io/ioutil"
  "encoding/json"
  "log"
  "fmt"
  "bytes"
  "strconv"
)

type Update struct {
	UpdateId int `json:"update_id"`
	Message Message `json:"message"`
}

type Message struct {
	Chat Chat `json:"chat"`
	Text string `json:"text"`
	Reply_markup ReplyKeyboardMarkup `json:"reply_markup"`
}

type BotMessage struct {
	ChatID int `json:"chat_id"`
	Text string `json:"text"`
	Reply_markup ReplyKeyboardMarkup `json:"reply_markup"`
}

type Chat struct {
	ChatID int `json:"id"`
}

type RestResponse struct {
	Result []Update `json:"result"`
}

type KeyboardButton struct {
	Text string `json:"text"`
}

type ReplyKeyboardMarkup struct {
	Keyboard [][]KeyboardButton `json:"keyboard"`
}

func getUpdates(botUrl string, offset int)([]Update, error){
	resp, err := http.Get(botUrl+"/getUpdates"+"?offset="+strconv.Itoa(offset))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil,err
	}
	var restResponse RestResponse
	err = json.Unmarshal(body, &restResponse)
	if err != nil {
		return nil, err
	}
	return restResponse.Result, nil
}

func respond(botUrl string, update Update, reply_markup ReplyKeyboardMarkup) (error){
	var botMessage BotMessage
	botMessage.ChatID = update.Message.Chat.ChatID
	botMessage.Text = update.Message.Text
	botMessage.Reply_markup = reply_markup

	buf, err := json.Marshal(botMessage)
	if err != nil {
		return err
	}
	_, err = http.Post(botUrl + "/sendMessage","application/json", bytes.NewBuffer(buf))
	if err != nil {
		return err
	}
	return nil
}

func mainMenu()(ReplyKeyboardMarkup){
  var button1, button2, button3, button4, button5, button6, button7, button8, button9, button10, button11, button12, button13 KeyboardButton
button1.Text = text1
button2.Text = text2
button3.Text = text3
button4.Text = text4
button5.Text = text5
button6.Text = text6
button7.Text = text7
button8.Text = text8
button9.Text = text9
button10.Text = text10
button11.Text = text11
button12.Text = text12
button13.Text = text13
  var but ReplyKeyboardMarkup
  but.Keyboard = [][]KeyboardButton{{button1}, {button2}, {button3}, {button4}, {button5}, {button6}, {button7}, {button8}, {button9}, {button10}, {button11}, {button12}, {button13}} 
return but
}

const text1 = "Текущая эпидемиологическая ситуация в РК"
const text2 = "Онлайн трансляция/брифинги госорганов"
const text3 = "Въезд/Выезд в/из РК"
const text4 = "Меры профилактики"
const text5 = "Куда обратиться в вашем городе, если есть симптомы"
const text6 = "Акты органов власти по коронавирусу"
const text7 = "Соблюдение трудовых прав"
const text8 = "Режим работы организаций в вашем городе"
const text9 = "Аптеки вашего города, в которых есть маски и антисептики"
const text10 = "Возврат авиа/жд билетов"
const text11 = "Контакты консульств РК за рубежом"
const text12 = "Актуальная информация для школьников и студентов РК"
const text13 = "Информация для обладателей стипендии Болашак"
func main() {
  // подключаемся к боту с помощью токена
  botToken := "1086353471:AAHByBefKoJ3aKFbDBLax5KoemW7fDSK0TQ"
  botApi := "https://api.telegram.org/bot"
  botUrl := botApi + botToken
  offset := 0
  but := mainMenu()

  for ;; {
  	updates, err := getUpdates(botUrl, offset)
  	if err != nil {
  		log.Println("Smth went wrong:", err.Error())
  	}
  	for _, update := range updates {
  		if update.Message.Text == text1 {
  			update.Message.Text = "Текущая эпидемиологическая ситуация в РК обновляется каждый день на сайте: https://web.facebook.com/MinzdravRK/?_rdc=1&_rdr"
  			err = respond(botUrl, update, but)

  			offset = update.UpdateId+1	
  			continue
  		}
  		if update.Message.Text == text2 {
  			update.Message.Text = "Здесь будет ссылка на онлайн брифинги. На данный момент трансляция не активна"
  			err = respond(botUrl, update, but)

  			offset = update.UpdateId+1	
  			continue
  		}
  		err = respond(botUrl, update, but)
  		offset = update.UpdateId+1
  	}
  	fmt.Println(updates)
  }

}