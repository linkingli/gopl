

/*curl -X POST \
http://10.125.30.68:46138/templates \
-H 'Accept: application/json' \
-H 'Cache-Control: no-cache' \
-H 'Content-Type: application/json' \
-H 'Postman-Token: 7c807cd1-26da-b7c9-85f4-5ab2d41cb4e3' \
-d '{
"template": {
"name": "测试模板",
"type": [
"EMAIL"
],
"content": "完成",
"enabled": true
}
}'
*/
/*curl -X POST \
http://10.125.30.68:46138/to \
-H 'Accept: application/json' \
-H 'Cache-Control: no-cache' \
-H 'Content-Type: application/json' \
-H 'Postman-Token: 35238030-a646-a794-62ae-883ac8087305' \
-d '{
"notification": {
"template_id": "7639392b-0242-470f-a7d4-5b28908dd905",
"transports": [
{
"transport": "EMAIL",
"to": "@.xx.com"
}

],
"parameters": {
"workissue_uuid": "123"
}
}
}
'*/




type CreateTemplate struct {
	Template   Template    `json:"template"`
}

type Template struct {
	Name string `json:"name"`
	Content string `json:"content"`
	Enabled bool `json:"enabled,omitempty"`
	Type [] string `json:"type"`
	Uuid  string  `json:"uuid,omitempty"`
}


type  SendEmail struct {
	Notification Notification `json:"notification"`
}
type  Notification struct {
	TemplateId string `json:"template_id"`
	TransportsArry []  Transports    `json:"transports"`
	Parameters  Parameters    `json:"parameters"`
}
type Transports struct {
	Transport string `json:"transport"`
	To  string    `json:"to"`
}
type Parameters struct {
	WorkissueUuid string `json:"workissue_uuid"`

}

type EmailRecive struct {
	Notificationhandles [] Notificationhandles `json:"notification_handles"`
}
type Notificationhandles struct {
	Uuid string `json:"uuid"`
	Transport  string    `json:"transport"`
}

func (t Tag)List(c *gin.Context)  {
	/*
	发送邮件
	 */
/*	var sendEmail SendEmail
	var parameters Parameters
	var transports Transports
	var notification Notification
	parameters.WorkissueUuid="123"
	transports.To="xxx@h3c.com"
	transports.Transport="SMS"

	notification.TransportsArry = append(notification.TransportsArry,transports)
	notification.Parameters=parameters
	notification.TemplateId="e4caea21-32a8-4a5b-a3fd-446060c3b3a8"
	sendEmail.Notification=notification
	marshal, _ := json.Marshal(sendEmail)
	fmt.Println(string(marshal))
	req := bytes.NewBuffer([]byte(marshal))

	client := &http.Client{}
	request, _ := http.NewRequest("POST", "http://10.125.30.68:46138/to", req)
	request.Header.Set("Content-type", "application/json")
	response, _ := client.Do(request)
	body, _ := ioutil.ReadAll(response.Body)
	if response.StatusCode == 202 {
		fmt.Println("发送邮件短信成功"+string(body))
	}
	var emailRecive EmailRecive
	json.Unmarshal(body, &emailRecive)
	var Uuid string
	notificationhandles := emailRecive.Notificationhandles
	for _, value :=range notificationhandles{
		if value.Uuid!="" {
			Uuid = value.Uuid
		}
	}
	fmt.Println("邮件短信id为"+Uuid)
*/
	/*
	 创建模板
	 */
	var createTemplate CreateTemplate
	var template Template
	template.Content="完成"
	template.Enabled=true
	template.Name="测试模板"
	template.Type = append(template.Type, "SMS")
	createTemplate.Template=template
	createMarshal, _ := json.Marshal(createTemplate)
	fmt.Println("创建邮件模板"+string(createMarshal))
	reqCreateTemplate := bytes.NewBuffer([]byte(createMarshal))

	clientcreateTemplate := &http.Client{}
	requestCreateTemplate, _ := http.NewRequest("POST", "http://10.125.30.68:46138/templates", reqCreateTemplate)
	requestCreateTemplate.Header.Set("Accept", "application/json")
	requestCreateTemplate.Header.Set("Content-type", "application/json")
	responseCreateTemplate, _ := clientcreateTemplate.Do(requestCreateTemplate)
	bodyCreateTemplate, _ := ioutil.ReadAll(responseCreateTemplate.Body)
	if responseCreateTemplate.StatusCode == 201 {
		fmt.Println("模板添加成功"+string(bodyCreateTemplate))
	}

	var resJsonCreateTemplate CreateTemplate
	json.Unmarshal(bodyCreateTemplate, &resJsonCreateTemplate)
	uuid := resJsonCreateTemplate.Template.Uuid
	fmt.Println("模板uuid",uuid)
	//todo 发送 email

}
