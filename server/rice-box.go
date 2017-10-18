package server

import (
	"github.com/GeertJohan/go.rice/embedded"
	"time"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    "add-device-success.gohtml",
		FileModTime: time.Unix(1508347479, 0),
		Content:     string("<html lang=\"en\" >\r\n    <head>\r\n        <title>Remote Wake/Sleep-On-LAN</title>\r\n        <meta http-equiv=\"Content-Type\" content=\"text/html;charset=utf-8\">\r\n        <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\r\n        <meta name=\"description\" content=\"A utility for remotely waking/sleeping a Windows computer via a Raspberry Pi\">\r\n        <meta name=\"author\" content=\"Carlo Maiorano\">\r\n        <meta http-equiv=\"cache-control\" content=\"max-age=0\" />\r\n        <meta http-equiv=\"cache-control\" content=\"no-cache\" />\r\n        <meta http-equiv=\"expires\" content=\"0\" />\r\n        <meta http-equiv=\"expires\" content=\"Tue, 01 Jan 1980 1:00:00 GMT\" />\r\n        <meta http-equiv=\"pragma\" content=\"no-cache\" />\r\n        <style type=\"text/css\">\r\n            body {\r\n                padding-top: 40px !important;\r\n                padding-bottom: 40px;\r\n                background-color: #f5f5f5;\r\n            }\r\n            .jumbotron {\r\n                max-width: 600px;\r\n                padding: 19px 29px 29px;\r\n                margin: 0 auto 20px;\r\n                background-color: #fff;\r\n                border: 1px solid #e5e5e5;\r\n                -webkit-border-radius: 5px;\r\n                -moz-border-radius: 5px;\r\n                border-radius: 5px;\r\n                -webkit-box-shadow: 0 1px 2px rgba(0,0,0,.05);\r\n                -moz-box-shadow: 0 1px 2px rgba(0,0,0,.05);\r\n                box-shadow: 0 1px 2px rgba(0,0,0,.05);\r\n            }\r\n        </style>\r\n        <link rel=\"stylesheet\" href=\"https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-beta/css/bootstrap.min.css\" integrity=\"sha384-/Y6pD6FV/Vv2HJnA6t+vslU6fwYXjCFtcEpHbNJ0lyAFsXTsjBbfaDjzALeQsN6M\" crossorigin=\"anonymous\">\r\n        <script src=\"https://code.jquery.com/jquery-3.2.1.slim.min.js\" integrity=\"sha384-KJ3o2DKtIkvYIK3UENzmM7KCkRr/rE9/Qpg6aAZGJwFDMVNA/GpGFF93hXpG5KkN\" crossorigin=\"anonymous\"></script>\r\n        <script src=\"https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.11.0/umd/popper.min.js\" integrity=\"sha384-b/U6ypiBEHpOf/4+1nzFpr53nxSS+GLCkfwBdFNTxtclqqenISfwAzpKaMNFNmj4\" crossorigin=\"anonymous\"></script>\r\n        <script src=\"https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-beta/js/bootstrap.min.js\" integrity=\"sha384-h0AbiXch4ZDo7tp9hKZ4TsHbi047NrKGLO3SEJAg45jXxnGIfYzk4Si90RDIqNm1\" crossorigin=\"anonymous\"></script>\r\n    </head>\r\n    <body>\r\n        <nav class=\"navbar navbar-default\">\r\n            <div class=\"container-fluid\">\r\n            <div class=\"navbar-header\">\r\n                <a class=\"navbar-brand\" href=\"/\">WolServer</a>\r\n            </div>\r\n            <ul class=\"nav navbar-nav\">\r\n                <li><a href=\"/\">Index</a></li>\r\n                <li class=\"active\"><a href=\"/devices\">Devices</a></li>\r\n                <li><a href=\"/config\">Configuration</a></li>\r\n            </ul>\r\n            </div>\r\n        </nav>\r\n        <div class = \"container-fluid\">\r\n            <div class = \"jumbotron\">\r\n                {{ with .}}     \r\n                    <h2>Added device {{ .Name }} successfully!!!</h2>\r\n                    <h4>Using Mac: {{ .Device.Mac }} with IP: {{ .Device.IP }}</h4>\r\n                {{ end }}\r\n                <h4>Now return to <a href=\"/\">home</a> in order to wake up your device or <a href=\"/devices\">add</a> a new one</h4>\r\n            </div>\r\n        </div>\r\n    </body>\r\n</html>"),
	}
	file3 := &embedded.EmbeddedFile{
		Filename:    "add-device.gohtml",
		FileModTime: time.Unix(1508347479, 0),
		Content:     string("<html lang=\"en\" >\r\n        <head>\r\n          <title>Remote Wake/Sleep-On-LAN</title>\r\n          <meta http-equiv=\"Content-Type\" content=\"text/html;charset=utf-8\">\r\n          <meta name=\"description\" content=\"A utility for remotely waking/sleeping a Windows computer via a Raspberry Pi\">\r\n          <meta name=\"author\" content=\"Carlo Maiorano\">\r\n          <meta http-equiv=\"cache-control\" content=\"max-age=0\" />\r\n          <meta http-equiv=\"cache-control\" content=\"no-cache\" />\r\n          <meta http-equiv=\"expires\" content=\"0\" />\r\n          <meta http-equiv=\"expires\" content=\"Tue, 01 Jan 1980 1:00:00 GMT\" />\r\n          <meta http-equiv=\"pragma\" content=\"no-cache\" />\r\n          <meta charset=\"utf-8\">\r\n          <meta name=\"viewport\" content=\"width=device-width, initial-scale=1, shrink-to-fit=no\">\r\n          <style type=\"text/css\">\r\n            body {\r\n                padding-top: 40px !important;\r\n                padding-bottom: 40px;\r\n                background-color: #f5f5f5;\r\n            }\r\n\r\n            .form-horizontal {\r\n                max-width: 600px;\r\n                padding: 19px 29px 29px;\r\n                margin: 0 auto 20px;\r\n                background-color: #fff;\r\n                border: 1px solid #e5e5e5;\r\n                -webkit-border-radius: 5px;\r\n                -moz-border-radius: 5px;\r\n                border-radius: 5px;\r\n                -webkit-box-shadow: 0 1px 2px rgba(0,0,0,.05);\r\n                -moz-box-shadow: 0 1px 2px rgba(0,0,0,.05);\r\n                box-shadow: 0 1px 2px rgba(0,0,0,.05);\r\n            }\r\n            .form-horizontal .form-horizontal-heading,\r\n            .form-horizontal .checkbox {\r\n                margin-bottom: 10px;\r\n            }\r\n            .form-horizontal input[type=\"text\"],\r\n            .form-horizontal input[type=\"password\"] {\r\n                font-size: 16px;\r\n                height: auto;\r\n                margin-bottom: 15px;\r\n                padding: 7px 9px;\r\n            }\r\n          </style>\r\n          <link rel=\"stylesheet\" href=\"https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-beta/css/bootstrap.min.css\" integrity=\"sha384-/Y6pD6FV/Vv2HJnA6t+vslU6fwYXjCFtcEpHbNJ0lyAFsXTsjBbfaDjzALeQsN6M\" crossorigin=\"anonymous\">\r\n          <script src=\"https://code.jquery.com/jquery-3.2.1.slim.min.js\" integrity=\"sha384-KJ3o2DKtIkvYIK3UENzmM7KCkRr/rE9/Qpg6aAZGJwFDMVNA/GpGFF93hXpG5KkN\" crossorigin=\"anonymous\"></script>\r\n          <script src=\"https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.11.0/umd/popper.min.js\" integrity=\"sha384-b/U6ypiBEHpOf/4+1nzFpr53nxSS+GLCkfwBdFNTxtclqqenISfwAzpKaMNFNmj4\" crossorigin=\"anonymous\"></script>\r\n          <script src=\"https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-beta/js/bootstrap.min.js\" integrity=\"sha384-h0AbiXch4ZDo7tp9hKZ4TsHbi047NrKGLO3SEJAg45jXxnGIfYzk4Si90RDIqNm1\" crossorigin=\"anonymous\"></script>\r\n        </head>\r\n        <body>\r\n            <nav class=\"navbar navbar-default\">\r\n                <div class=\"container-fluid\">\r\n                <div class=\"navbar-header\">\r\n                    <a class=\"navbar-brand\" href=\"/\">WolServer</a>\r\n                </div>\r\n                    <ul class=\"nav navbar-nav\">\r\n                        <li class=\"active\"><a href=\"/\">Index</a></li>\r\n                        <li><a href=\"/devices\">Devices</a></li>\r\n                        <li><a href=\"/config\">Configuration</a></li>\r\n                    </ul>\r\n                </div>\r\n            </nav>\r\n            <div class=\"container-fluid\">\r\n                <form class=\"form-horizontal\" action=\"/devices\" method=\"POST\">\r\n                    <h2 class=\"form-horizontal-heading\">Add a new device</h2>\r\n                    <p class=\"form-horizontal-heading\">Please fill the form with device data</p>\r\n                    <div class=\"form-group\">\r\n                        <label class=\"col-sm-3 control-label\">Password:</label>\r\n                        <div class=\"col-sm-10\">\r\n                            <input type=\"password\" class=\"form-control\" name=\"password\" id=\"validationPass\" placeholder=\"Password\">\r\n                        </div>\r\n                    </div>\r\n                    <div class=\"form-group\">\r\n                        <label class=\"col-sm-3 control-label\">Mac Address:</label>\r\n                        <div class=\"col-sm-10\">\r\n                            <input type=\"text\" class=\"form-control\" name=\"macAddr\" id=\"configMacAddr\" placeholder=\"Mac Address\">\r\n                        </div>\r\n                    </div>\r\n                    <div class=\"form-group\">\r\n                        <label class=\"col-sm-3 control-label\">IP Address:</label>\r\n                        <div class=\"col-sm-10\">\r\n                            <input type=\"text\" class=\"form-control\" name=\"ipAddr\" id=\"configMacAddr\" placeholder=\"IP Address\">\r\n                        </div>\r\n                    </div>\r\n                    <div class=\"form-group\">\r\n                        <label class=\"col-sm-3 control-label\">Alias:</label>\r\n                        <div class=\"col-sm-10\">\r\n                            <input type=\"text\" class=\"form-control\" name=\"alias\" id=\"configAlias\" placeholder=\"Alias\">\r\n                        </div>\r\n                    </div>\r\n                    <div class=\"form-group\">\r\n                        <div class=\"col-sm-offset-2 col-sm-10\">\r\n                            <button type=\"submit\" class=\"btn btn-large btn-primary\">Add</button>\r\n                        </div>\r\n                    </div>\r\n                </form>\r\n            </div>\r\n        </body>\r\n</html>"),
	}
	file4 := &embedded.EmbeddedFile{
		Filename:    "config-success.html",
		FileModTime: time.Unix(1508347479, 0),
		Content:     string("<html lang=\"en\" >\r\n    <head>\r\n      <title>Remote Wake/Sleep-On-LAN</title>\r\n      <meta http-equiv=\"Content-Type\" content=\"text/html;charset=utf-8\">\r\n      <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\r\n      <meta name=\"description\" content=\"A utility for remotely waking/sleeping a Windows computer via a Raspberry Pi\">\r\n      <meta name=\"author\" content=\"Carlo Maiorano\">\r\n      <meta http-equiv=\"cache-control\" content=\"max-age=0\" />\r\n      <meta http-equiv=\"cache-control\" content=\"no-cache\" />\r\n      <meta http-equiv=\"expires\" content=\"0\" />\r\n      <meta http-equiv=\"expires\" content=\"Tue, 01 Jan 1980 1:00:00 GMT\" />\r\n      <meta http-equiv=\"pragma\" content=\"no-cache\" />\r\n      <style type=\"text/css\">\r\n        body {\r\n            padding-top: 40px !important;\r\n            padding-bottom: 40px;\r\n            background-color: #f5f5f5;\r\n        }\r\n\r\n        .jumbotron {\r\n            max-width: 600px;\r\n            padding: 19px 29px 29px;\r\n            margin: 0 auto 20px;\r\n            background-color: #fff;\r\n            border: 1px solid #e5e5e5;\r\n            -webkit-border-radius: 5px;\r\n            -moz-border-radius: 5px;\r\n            border-radius: 5px;\r\n            -webkit-box-shadow: 0 1px 2px rgba(0,0,0,.05);\r\n            -moz-box-shadow: 0 1px 2px rgba(0,0,0,.05);\r\n            box-shadow: 0 1px 2px rgba(0,0,0,.05);\r\n        }\r\n      </style>\r\n      <link rel=\"stylesheet\" href=\"https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-beta/css/bootstrap.min.css\" integrity=\"sha384-/Y6pD6FV/Vv2HJnA6t+vslU6fwYXjCFtcEpHbNJ0lyAFsXTsjBbfaDjzALeQsN6M\" crossorigin=\"anonymous\">\r\n      <script src=\"https://code.jquery.com/jquery-3.2.1.slim.min.js\" integrity=\"sha384-KJ3o2DKtIkvYIK3UENzmM7KCkRr/rE9/Qpg6aAZGJwFDMVNA/GpGFF93hXpG5KkN\" crossorigin=\"anonymous\"></script>\r\n      <script src=\"https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.11.0/umd/popper.min.js\" integrity=\"sha384-b/U6ypiBEHpOf/4+1nzFpr53nxSS+GLCkfwBdFNTxtclqqenISfwAzpKaMNFNmj4\" crossorigin=\"anonymous\"></script>\r\n      <script src=\"https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-beta/js/bootstrap.min.js\" integrity=\"sha384-h0AbiXch4ZDo7tp9hKZ4TsHbi047NrKGLO3SEJAg45jXxnGIfYzk4Si90RDIqNm1\" crossorigin=\"anonymous\"></script>\r\n    </head>\r\n    <body>\r\n        <nav class=\"navbar navbar-default\">\r\n            <div class=\"container-fluid\">\r\n                <div class=\"navbar-header\">\r\n                <a class=\"navbar-brand\" href=\"/\">WolServer</a>\r\n                </div>\r\n                <ul class=\"nav navbar-nav\">\r\n                <li><a href=\"/\">Index</a></li>\r\n                <li><a href=\"/devices\">Devices</a></li>\r\n                <li class=\"active\"><a href=\"/config\">Configuration</a></li>\r\n                </ul>\r\n            </div>\r\n        </nav>\r\n        <div class=\"container-fluid\">\r\n            <div class = \"jumbotron\">\r\n                <h2>Configuration Successfull</h2>\r\n                <h4>Now add some <a href=\"/devices\">devices</a></h4>\r\n            </div>\r\n        </div>\r\n    </body>\r\n</html>"),
	}
	file5 := &embedded.EmbeddedFile{
		Filename:    "config.html",
		FileModTime: time.Unix(1508347479, 0),
		Content:     string("<html lang=\"en\" >\r\n    <head>\r\n      <title>Remote Wake/Sleep-On-LAN</title>\r\n      <meta http-equiv=\"Content-Type\" content=\"text/html;charset=utf-8\">\r\n      <meta name=\"description\" content=\"A utility for remotely waking/sleeping a Windows computer via a Raspberry Pi\">\r\n      <meta name=\"author\" content=\"Carlo Maiorano\">\r\n      <meta http-equiv=\"cache-control\" content=\"max-age=0\" />\r\n      <meta http-equiv=\"cache-control\" content=\"no-cache\" />\r\n      <meta http-equiv=\"expires\" content=\"0\" />\r\n      <meta http-equiv=\"expires\" content=\"Tue, 01 Jan 1980 1:00:00 GMT\" />\r\n      <meta http-equiv=\"pragma\" content=\"no-cache\" />\r\n      <style type=\"text/css\">\r\n        body {\r\n            padding-top: 40px !important;\r\n            padding-bottom: 40px;\r\n            background-color: #f5f5f5;\r\n        }\r\n\r\n        .form-horizontal {\r\n            max-width: 600px;\r\n            padding: 19px 29px 29px;\r\n            margin: 0 auto 20px;\r\n            background-color: #fff;\r\n            border: 1px solid #e5e5e5;\r\n            -webkit-border-radius: 5px;\r\n            -moz-border-radius: 5px;\r\n            border-radius: 5px;\r\n            -webkit-box-shadow: 0 1px 2px rgba(0,0,0,.05);\r\n            -moz-box-shadow: 0 1px 2px rgba(0,0,0,.05);\r\n            box-shadow: 0 1px 2px rgba(0,0,0,.05);\r\n        }\r\n        .form-horizontal .form-horizontal-heading,\r\n        .form-horizontal .checkbox {\r\n            margin-bottom: 10px;\r\n        }\r\n        .form-horizontal input[type=\"text\"],\r\n        .form-horizontal input[type=\"password\"] {\r\n            font-size: 16px;\r\n            height: auto;\r\n            margin-bottom: 15px;\r\n            padding: 7px 9px;\r\n        }\r\n      </style>\r\n      <link rel=\"stylesheet\" href=\"https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-beta/css/bootstrap.min.css\" integrity=\"sha384-/Y6pD6FV/Vv2HJnA6t+vslU6fwYXjCFtcEpHbNJ0lyAFsXTsjBbfaDjzALeQsN6M\" crossorigin=\"anonymous\">\r\n      <script src=\"https://code.jquery.com/jquery-3.2.1.slim.min.js\" integrity=\"sha384-KJ3o2DKtIkvYIK3UENzmM7KCkRr/rE9/Qpg6aAZGJwFDMVNA/GpGFF93hXpG5KkN\" crossorigin=\"anonymous\"></script>\r\n      <script src=\"https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.11.0/umd/popper.min.js\" integrity=\"sha384-b/U6ypiBEHpOf/4+1nzFpr53nxSS+GLCkfwBdFNTxtclqqenISfwAzpKaMNFNmj4\" crossorigin=\"anonymous\"></script>\r\n      <script src=\"https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-beta/js/bootstrap.min.js\" integrity=\"sha384-h0AbiXch4ZDo7tp9hKZ4TsHbi047NrKGLO3SEJAg45jXxnGIfYzk4Si90RDIqNm1\" crossorigin=\"anonymous\"></script>\r\n      <link href=\"style/bootstrap/css/bootstrap.css\" rel=\"stylesheet\">\r\n      <link href=\"style/bootstrap/css/bootstrap-responsive.css\" rel=\"stylesheet\">\r\n    </head>\r\n    <body>\r\n        <nav class=\"navbar navbar-default\">\r\n            <div class=\"container-fluid\">\r\n                <div class=\"navbar-header\">\r\n                <a class=\"navbar-brand\" href=\"/\">WolServer</a>\r\n                </div>\r\n                <ul class=\"nav navbar-nav\">\r\n                <li><a href=\"/\">Index</a></li>\r\n                <li><a href=\"/devices\">Devices</a></li>\r\n                <li class=\"active\"><a href=\"/config\">Configuration</a></li>\r\n                </ul>\r\n            </div>\r\n        </nav>\r\n        <div class=\"container-fluid\">\r\n            <form class=\"form-horizontal\" action=\"/config\" method=\"POST\">\r\n                <h2 class=\"form-horizontal-heading\">Configuration</h2>\r\n                <p class=\"form-horizontal-heading\">Please insert initial password in order to configure the application</p>\r\n                <div class=\"form-group\">\r\n                    <label class=\"col-sm-5 control-label\">Configuration Password:</label>\r\n                    <div class=\"col-sm-10\">\r\n                        <input type=\"password\" class=\"form-control\" name=\"password\" id=\"configPassword\" placeholder=\"Password\">\r\n                    </div>\r\n                </div>\r\n                <div class=\"form-group\">\r\n                    <div class=\"col-sm-offset-2 col-sm-10\">\r\n                        <button type=\"submit\" class=\"btn btn-large btn-primary\">Add</button>\r\n                    </div>\r\n                </div>\r\n            </form>\r\n        </div>\r\n    </body>\r\n</html>"),
	}
	file6 := &embedded.EmbeddedFile{
		Filename:    "error.gohtml",
		FileModTime: time.Unix(1508347479, 0),
		Content:     string("<html lang=\"en\" >\r\n    <head>\r\n        <title>Remote Wake/Sleep-On-LAN</title>\r\n        <meta http-equiv=\"Content-Type\" content=\"text/html;charset=utf-8\">\r\n        <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\r\n        <meta name=\"description\" content=\"A utility for remotely waking/sleeping a Windows computer via a Raspberry Pi\">\r\n        <meta name=\"author\" content=\"Carlo Maiorano\">\r\n        <meta http-equiv=\"cache-control\" content=\"max-age=0\" />\r\n        <meta http-equiv=\"cache-control\" content=\"no-cache\" />\r\n        <meta http-equiv=\"expires\" content=\"0\" />\r\n        <meta http-equiv=\"expires\" content=\"Tue, 01 Jan 1980 1:00:00 GMT\" />\r\n        <meta http-equiv=\"pragma\" content=\"no-cache\" />\r\n        <style type=\"text/css\">\r\n            body {\r\n                padding-top: 40px !important;\r\n                padding-bottom: 40px;\r\n                background-color: #f5f5f5;\r\n            }\r\n        </style>\r\n        <link rel=\"stylesheet\" href=\"https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-beta/css/bootstrap.min.css\" integrity=\"sha384-/Y6pD6FV/Vv2HJnA6t+vslU6fwYXjCFtcEpHbNJ0lyAFsXTsjBbfaDjzALeQsN6M\" crossorigin=\"anonymous\">\r\n        <script src=\"https://code.jquery.com/jquery-3.2.1.slim.min.js\" integrity=\"sha384-KJ3o2DKtIkvYIK3UENzmM7KCkRr/rE9/Qpg6aAZGJwFDMVNA/GpGFF93hXpG5KkN\" crossorigin=\"anonymous\"></script>\r\n        <script src=\"https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.11.0/umd/popper.min.js\" integrity=\"sha384-b/U6ypiBEHpOf/4+1nzFpr53nxSS+GLCkfwBdFNTxtclqqenISfwAzpKaMNFNmj4\" crossorigin=\"anonymous\"></script>\r\n        <script src=\"https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-beta/js/bootstrap.min.js\" integrity=\"sha384-h0AbiXch4ZDo7tp9hKZ4TsHbi047NrKGLO3SEJAg45jXxnGIfYzk4Si90RDIqNm1\" crossorigin=\"anonymous\"></script>\r\n    </head>\r\n    <body>\r\n        <nav class=\"navbar navbar-default\">\r\n            <div class=\"container-fluid\">\r\n            <div class=\"navbar-header\">\r\n                <a class=\"navbar-brand\" href=\"/\">WolServer</a>\r\n            </div>\r\n                <ul class=\"nav navbar-nav\">\r\n                    <li><a href=\"/\">Index</a></li>\r\n                    <li><a href=\"/devices\">Devices</a></li>\r\n                    <li><a href=\"/config\">Configuration</a></li>\r\n                </ul>\r\n            </div>\r\n        </nav>\r\n        <div class=\"container-fluid\">\r\n            <div class=\"jumbotron\">\r\n                <h2>Error!!!</h2>\r\n                {{ with . }}\r\n                    <h4>Got error: {{.Message}}</h4>\r\n                {{ end }}\r\n            </div>\r\n        <div>\r\n    </body>\r\n</html>"),
	}
	file7 := &embedded.EmbeddedFile{
		Filename:    "index.gohtml",
		FileModTime: time.Unix(1508349051, 0),
		Content:     string("<html lang=\"en\" >\r\n    <head>\r\n      <title>Remote Wake/Sleep-On-LAN</title>\r\n      <meta http-equiv=\"Content-Type\" content=\"text/html;charset=utf-8\">\r\n      <meta name=\"description\" content=\"A utility for remotely waking/sleeping a Windows computer via a Raspberry Pi\">\r\n      <meta name=\"author\" content=\"Carlo Maiorano\">\r\n      <link href=\"style/bootstrap/css/bootstrap.css\" rel=\"stylesheet\">\r\n      <link href=\"style/bootstrap/css/bootstrap-responsive.css\" rel=\"stylesheet\">\r\n      <meta http-equiv=\"cache-control\" content=\"max-age=0\" />\r\n      <meta http-equiv=\"cache-control\" content=\"no-cache\" />\r\n      <meta http-equiv=\"expires\" content=\"0\" />\r\n      <meta http-equiv=\"expires\" content=\"Tue, 01 Jan 1980 1:00:00 GMT\" />\r\n      <meta http-equiv=\"pragma\" content=\"no-cache\" />\r\n      <meta charset=\"utf-8\">\r\n      <meta name=\"viewport\" content=\"width=device-width, initial-scale=1, shrink-to-fit=no\">\r\n      <style type=\"text/css\">\r\n            body {\r\n                padding-top: 40px !important;\r\n                padding-bottom: 40px;\r\n                background-color: #f5f5f5;\r\n            }\r\n\r\n            .form-horizontal {\r\n                max-width: 600px;\r\n                padding: 19px 29px 29px;\r\n                margin: 0 auto 20px;\r\n                background-color: #fff;\r\n                border: 1px solid #e5e5e5;\r\n                -webkit-border-radius: 5px;\r\n                -moz-border-radius: 5px;\r\n                border-radius: 5px;\r\n                -webkit-box-shadow: 0 1px 2px rgba(0,0,0,.05);\r\n                -moz-box-shadow: 0 1px 2px rgba(0,0,0,.05);\r\n                box-shadow: 0 1px 2px rgba(0,0,0,.05);\r\n            }\r\n            .form-horizontal .form-horizontal-heading,\r\n            .form-horizontal .checkbox {\r\n                margin-bottom: 10px;\r\n            }\r\n            .form-horizontal input[type=\"text\"],\r\n            .form-horizontal input[type=\"password\"] {\r\n                font-size: 16px;\r\n                height: auto;\r\n                margin-bottom: 15px;\r\n                padding: 7px 9px;\r\n            }\r\n      </style>\r\n      <link rel=\"stylesheet\" href=\"https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-beta/css/bootstrap.min.css\" integrity=\"sha384-/Y6pD6FV/Vv2HJnA6t+vslU6fwYXjCFtcEpHbNJ0lyAFsXTsjBbfaDjzALeQsN6M\" crossorigin=\"anonymous\">\r\n      <script src=\"https://code.jquery.com/jquery-3.2.1.slim.min.js\" integrity=\"sha384-KJ3o2DKtIkvYIK3UENzmM7KCkRr/rE9/Qpg6aAZGJwFDMVNA/GpGFF93hXpG5KkN\" crossorigin=\"anonymous\"></script>\r\n      <script src=\"https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.11.0/umd/popper.min.js\" integrity=\"sha384-b/U6ypiBEHpOf/4+1nzFpr53nxSS+GLCkfwBdFNTxtclqqenISfwAzpKaMNFNmj4\" crossorigin=\"anonymous\"></script>\r\n      <script src=\"https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-beta/js/bootstrap.min.js\" integrity=\"sha384-h0AbiXch4ZDo7tp9hKZ4TsHbi047NrKGLO3SEJAg45jXxnGIfYzk4Si90RDIqNm1\" crossorigin=\"anonymous\"></script>\r\n      <!-- HTML5 shim, for IE6-8 support of HTML5 elements -->\r\n      <!--[if lt IE 9]>\r\n        <script src=\"style/bootstrap/js/html5shiv.js\"></script>\r\n      <![endif]-->\r\n      <link rel=\"apple-touch-icon-precomposed\" sizes=\"144x144\" href=\"style/bootstrap/ico/apple-touch-icon-144-precomposed.png\">\r\n      <link rel=\"apple-touch-icon-precomposed\" sizes=\"114x114\" href=\"style/bootstrap/ico/apple-touch-icon-114-precomposed.png\">\r\n      <link rel=\"apple-touch-icon-precomposed\" sizes=\"72x72\" href=\"style/bootstrap/ico/apple-touch-icon-72-precomposed.png\">\r\n      <link rel=\"apple-touch-icon-precomposed\" href=\"style/bootstrap/ico/apple-touch-icon-57-precomposed.png\">\r\n      <link rel=\"shortcut icon\" href=\"style/bootstrap/ico/favicon.png\">\r\n    </head>\r\n    <body>\r\n      <nav class=\"navbar navbar-toggleable-md navbar-light bg-faded\">\r\n        <button class=\"navbar-toggler navbar-toggler-right\" type=\"button\" data-toggle=\"collapse\" data-target=\"#navbarSupportedContent\" aria-controls=\"navbarSupportedContent\" aria-expanded=\"false\" aria-label=\"Toggle navigation\">\r\n          <span class=\"navbar-toggler-icon\"></span>\r\n        </button>\r\n        <a class=\"navbar-brand\" href=\"/\">Wol</a>\r\n\r\n        <div class=\"collapse navbar-collapse\" id=\"navbarSupportedContent\">\r\n          <ul class=\"navbar-nav mr-auto\">\r\n            <li class=\"nav-item active\">\r\n              <a class=\"nav-link\" href=\"/\">Home <span class=\"sr-only\">(current)</span></a>\r\n            </li>\r\n            <li class=\"nav-item\">\r\n              <a class=\"nav-link\" href=\"/devices\">Devices</a>\r\n            </li>\r\n            <li class=\"nav-item\">\r\n              <a class=\"nav-link\" href=\"/config\">Configuration</a>\r\n            </li>\r\n          </ul>\r\n        </div>\r\n      </nav>\r\n      <div class=\"container-fluid\">\r\n        <form  class=\"form-horizontal\" action=\"/\" method=\"POST\">\r\n          <h2 class=\"form-horizontal-heading\">Select Target Device</h2>\r\n          <p class=\"form-horizontal-heading\">Select wake on lan target</p>\r\n          <div class=\"form-group\">\r\n            <label class=\"col-sm-3 control-label\">Select Device:</label>\r\n            <div class=\"col-sm-10\">\r\n              <select  class=\"form-control\" name=\"devices\">\r\n                {{ range $ }}\r\n                  <option value=\"{{ . }}\">{{ . }}</option>\r\n                {{ end }}\r\n              </select>\r\n            </div>\r\n          </div>\r\n          <div class=\"form-group\">\r\n            <label class=\"col-sm-3 control-label\">Password:</label>\r\n            <div class=\"col-sm-10\">\r\n              <input type=\"password\" class=\"form-control\" name=\"password\" id=\"wakeupPassword\" placeholder=\"Password\">\r\n            </div>\r\n          </div>\r\n           <div class=\"form-group\">\r\n                <div class=\"col-sm-offset-2 col-sm-10\">\r\n                    <button type=\"submit\" class=\"btn btn-large btn-primary\">Wake Up!</button>\r\n                </div>\r\n            </div>\r\n        </form>\r\n      </div>\r\n    </body>\r\n  </html>"),
	}
	file8 := &embedded.EmbeddedFile{
		Filename:    "report.gohtml",
		FileModTime: time.Unix(1508347479, 0),
		Content:     string("<html lang=\"en\" >\r\n    <head>\r\n      <title>Remote Wake/Sleep-On-LAN</title>\r\n      <meta http-equiv=\"Content-Type\" content=\"text/html;charset=utf-8\">\r\n      <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\r\n      <meta name=\"description\" content=\"A utility for remotely waking/sleeping a Windows computer via a Raspberry Pi\">\r\n      <meta name=\"author\" content=\"Carlo Maiorano\">\r\n      <link href=\"style/bootstrap/css/bootstrap.css\" rel=\"stylesheet\">\r\n      <link href=\"style/bootstrap/css/bootstrap-responsive.css\" rel=\"stylesheet\">\r\n      <meta http-equiv=\"cache-control\" content=\"max-age=0\" />\r\n      <meta http-equiv=\"cache-control\" content=\"no-cache\" />\r\n      <meta http-equiv=\"expires\" content=\"0\" />\r\n      <meta http-equiv=\"expires\" content=\"Tue, 01 Jan 1980 1:00:00 GMT\" />\r\n      <meta http-equiv=\"pragma\" content=\"no-cache\" />\r\n      <meta charset=\"utf-8\">\r\n      <meta name=\"viewport\" content=\"width=device-width, initial-scale=1, shrink-to-fit=no\">\r\n      <style type=\"text/css\">\r\n        body {\r\n          padding-top: 40px !important;\r\n          padding-bottom: 40px;\r\n          background-color: #f5f5f5;\r\n        }\r\n\r\n        .table-bordered {\r\n          max-width: 600px;\r\n          padding: 19px 29px 29px;\r\n          margin: 0 auto 20px;\r\n          background-color: #fff;\r\n          border: 1px solid #e5e5e5;\r\n          -webkit-border-radius: 5px;\r\n          -moz-border-radius: 5px;\r\n          border-radius: 5px;\r\n          -webkit-box-shadow: 0 1px 2px rgba(0,0,0,.05);\r\n          -moz-box-shadow: 0 1px 2px rgba(0,0,0,.05);\r\n          box-shadow: 0 1px 2px rgba(0,0,0,.05);\r\n        }\r\n        .table-bordered .table-bordered-heading,\r\n        .table-bordered .checkbox {\r\n            margin-bottom: 10px;\r\n        }\r\n      </style>\r\n      <link rel=\"stylesheet\" href=\"https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-beta/css/bootstrap.min.css\" integrity=\"sha384-/Y6pD6FV/Vv2HJnA6t+vslU6fwYXjCFtcEpHbNJ0lyAFsXTsjBbfaDjzALeQsN6M\" crossorigin=\"anonymous\">\r\n      <script src=\"https://code.jquery.com/jquery-3.2.1.slim.min.js\" integrity=\"sha384-KJ3o2DKtIkvYIK3UENzmM7KCkRr/rE9/Qpg6aAZGJwFDMVNA/GpGFF93hXpG5KkN\" crossorigin=\"anonymous\"></script>\r\n      <script src=\"https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.11.0/umd/popper.min.js\" integrity=\"sha384-b/U6ypiBEHpOf/4+1nzFpr53nxSS+GLCkfwBdFNTxtclqqenISfwAzpKaMNFNmj4\" crossorigin=\"anonymous\"></script>\r\n      <script src=\"https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-beta/js/bootstrap.min.js\" integrity=\"sha384-h0AbiXch4ZDo7tp9hKZ4TsHbi047NrKGLO3SEJAg45jXxnGIfYzk4Si90RDIqNm1\" crossorigin=\"anonymous\"></script>\r\n\r\n      <!-- HTML5 shim, for IE6-8 support of HTML5 elements -->\r\n      <!--[if lt IE 9]>\r\n        <script src=\"style/bootstrap/js/html5shiv.js\"></script>\r\n      <![endif]-->\r\n      <link rel=\"apple-touch-icon-precomposed\" sizes=\"144x144\" href=\"style/bootstrap/ico/apple-touch-icon-144-precomposed.png\">\r\n      <link rel=\"apple-touch-icon-precomposed\" sizes=\"114x114\" href=\"style/bootstrap/ico/apple-touch-icon-114-precomposed.png\">\r\n      <link rel=\"apple-touch-icon-precomposed\" sizes=\"72x72\" href=\"style/bootstrap/ico/apple-touch-icon-72-precomposed.png\">\r\n      <link rel=\"apple-touch-icon-precomposed\" href=\"style/bootstrap/ico/apple-touch-icon-57-precomposed.png\">\r\n      <link rel=\"shortcut icon\" href=\"style/bootstrap/ico/favicon.png\">\r\n    </head>\r\n    <body>\r\n      <nav class=\"navbar navbar-default\">\r\n        <div class=\"container-fluid\">\r\n          <div class=\"navbar-header\">\r\n            <a class=\"navbar-brand\" href=\"/\">WolServer</a>\r\n          </div>\r\n          <ul class=\"nav navbar-nav\">\r\n            <li class=\"active\"><a href=\"/\">Index</a></li>\r\n            <li><a href=\"/devices\">Devices</a></li>\r\n            <li><a href=\"/config\">Configuration</a></li>\r\n          </ul>\r\n        </div>\r\n      </nav>\r\n      <div class=\"container-fluid\">\r\n        <div class =\"jumbotron\">\r\n          <h1>Report</h1>\r\n          {{ with .}}\r\n          <h4>Summary of {{ .Alias }}:</h4>\r\n          <table class=\"table table-bordered\">\r\n            <thead>\r\n            <tr>\r\n              <th>Time</th>\r\n              <th>Result</th>\r\n            </tr>\r\n            </thead>\r\n            <tbody>\r\n              {{ range $time, $result := .Report }}\r\n                {{ if $result }}\r\n                  <tr class=\"success\">\r\n                    <td>{{ $time }}</td>\r\n                    <td>The device is alive</td>\r\n                  </tr>\r\n                {{ else  }}\r\n                  <tr class=\"danger\">\r\n                    <td>{{ $time }}</td>\r\n                    <td>The device is still sleeping</td>\r\n                  </tr>\r\n                {{ end }}\r\n              {{ end }}\r\n            </tbody>\r\n          </table>\r\n          {{ end }}\r\n        </div>\r\n      </div>\r\n    </body>\r\n  </html>"),
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   "",
		DirModTime: time.Unix(1508347479, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2, // "add-device-success.gohtml"
			file3, // "add-device.gohtml"
			file4, // "config-success.html"
			file5, // "config.html"
			file6, // "error.gohtml"
			file7, // "index.gohtml"
			file8, // "report.gohtml"

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`../templates/`, &embedded.EmbeddedBox{
		Name: `../templates/`,
		Time: time.Unix(1508347479, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir1,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"add-device-success.gohtml": file2,
			"add-device.gohtml":         file3,
			"config-success.html":       file4,
			"config.html":               file5,
			"error.gohtml":              file6,
			"index.gohtml":              file7,
			"report.gohtml":             file8,
		},
	})
}
