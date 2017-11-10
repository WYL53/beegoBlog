{{define "header"}}
<html>
<head>
  
    <link rel="shortcut" href="/sttic/img/favicon.png">
    <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css">

   <title>wyl的博客</title>
    <style >
    body {
        padding-top : 70px;
    }
    </style>
</head>

<body>
    <div class="bs-example bs-navbar-top-example" data-example-id="navbar-fixed-to-top">
            <div class="container">
               {{template "navbar" .}}
            </div>
    </div>
{{end}}