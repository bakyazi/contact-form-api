It is a simple endpoint for contact forms on your website.

It allows only rest `POST` method and accept following json

``` json

{
    "from": "email@address.com",
    "title": "title",
    "message": "text"
}

```
You can deploy this api on Google Cloud Run platform by clicking button below

[![Run on Google Cloud](https://deploy.cloud.run/button.svg)](https://deploy.cloud.run)


You should configure following enviroment variables

``` json
{
    "env": {
        "EMAIL":{
            "description": "Provide your Email Address",
            "required": true
        },
        "PASSWORD":{
            "description": "Provide your Password",
            "required": true
        },
        "NAME":{
            "description": "Provide your Name",
            "required": true
        }
    }
}
```

Enviroment variable `EMAIL` sends an e-mail which contains contact form payload to itself.

I use it for my [personal website](https://imberkay.com) and its [repo](https://github.com/bakyazi/nextjs-resume)

Here screenshot of contact form:

![contact-form](/assets/img1.jpg)

And screenshot of received e-mail:

![email](/assets/img3.jpg)


Please note that you should enable `Less secure app access` option of your gmail acount
![settings](/assets/img2.jpg)




Roadmap
--

 - [ ] API Rate Limiter
 - [ ] More fancy e-mail format via HTML
