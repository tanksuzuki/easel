package main

const TEMPLATE_CANVAS = `
<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  {{ template "sanitize" }}
  {{ template "style_canvas" }}
  {{ template "style_markdown" }}
</head>
<body>
  <div class="wrap">
    <div class="row">
      <div class="content">
        <h1>PROBLEM</h1>
        {{ .PROBLEM }}
      </div>
      <div class="col">
        <div class="content">
          <h1>SOLUTION</h1>
          {{ .SOLUTION }}
        </div>
        <div class="content">
          <h1>KEY METRICS</h1>
          {{ .KEYMETRICS }}
        </div>
      </div>
      <div class="content">
        <h1>UNIQUE VALUE PROPOSITION</h1>
        {{ .UNIQUEVALUEPROPOSITION }}
      </div>
      <div class="col">
        <div class="content">
          <h1>UNFAIR ADVANTAGE</h1>
          {{ .UNFAIRADVANTAGE }}
        </div>
        <div class="content">
          <h1>CHANNELS</h1>
          {{ .CHANNELS }}
        </div>
      </div>
      <div class="content">
        <h1>CUSTOMER SEGMENTS</h1>
        {{ .CUSTOMERSEGMENTS }}
      </div>
    </div>
    <div class="row">
      <div class="content">
        <h1>COST STRUCTURE</h1>
        {{ .COSTSTRUCTURE }}
      </div>
      <div class="content">
        <h1>REVENUE STREAMS</h1>
        {{ .REVENUESTREAMS }}
      </div>
    </div>
  </div>
  {{ with .livereload }}{{ . }}{{ end }}
</body>
</html>
`

const TEMPLATE_SANITIZE = `{{ define "sanitize" }}
<style>
/*! sanitize.css v4.1.0 | CC0 License | github.com/jonathantneal/sanitize.css */progress,sub,sup{vertical-align:baseline}button,hr,input{overflow:visible}[type=checkbox],[type=radio],legend{padding:0}[aria-disabled],html{cursor:default}article,aside,details,figcaption,figure,footer,header,main,menu,nav,section,summary{display:block}audio,canvas,progress,video{display:inline-block}audio:not([controls]){display:none;height:0}[hidden],template{display:none}*,::after,::before{background-repeat:no-repeat;box-sizing:inherit}::after,::before{text-decoration:inherit;vertical-align:inherit}html{box-sizing:border-box;font-family:sans-serif;line-height:1.5;-ms-text-size-adjust:100%;-webkit-text-size-adjust:100%}body{margin:0}h1{font-size:2em;margin:.67em 0}code,kbd,pre,samp{font-family:monospace,monospace;font-size:1em}small,sub,sup{font-size:83.3333%}hr{height:0}nav ol,nav ul{list-style:none}abbr[title]{border-bottom:1px dotted;text-decoration:none}b,strong{font-weight:bolder}dfn{font-style:italic}mark{background-color:#ff0;color:#000}sub,sup{line-height:0;position:relative}sub{bottom:-.25em}sup{top:-.5em}::-moz-selection{background-color:#b3d4fc;color:#000;text-shadow:none}::selection{background-color:#b3d4fc;color:#000;text-shadow:none}audio,canvas,iframe,img,svg,video{vertical-align:middle}img{border-style:none}svg{fill:currentColor}svg:not(:root){overflow:hidden}a{background-color:transparent;-webkit-text-decoration-skip:objects}a:hover{outline-width:0}table{border-collapse:collapse;border-spacing:0}button,input,select,textarea{background-color:transparent;border-style:none;color:inherit;font-size:1em;margin:0}button,select{text-transform:none}[type=reset],[type=submit],button,html [type=button]{-webkit-appearance:button}::-moz-focus-inner{border-style:none;padding:0}:-moz-focusring{outline:ButtonText dotted 1px}fieldset{border:1px solid silver;margin:0 2px;padding:.35em .625em .75em}legend{display:table;max-width:100%;white-space:normal}textarea{overflow:auto;resize:vertical}::-webkit-inner-spin-button,::-webkit-outer-spin-button{height:auto}[type=search]{-webkit-appearance:textfield;outline-offset:-2px}::-webkit-search-cancel-button,::-webkit-search-decoration{-webkit-appearance:none}::-webkit-input-placeholder{color:inherit;opacity:.54}::-webkit-file-upload-button{-webkit-appearance:button;font:inherit}[aria-busy=true]{cursor:progress}[aria-controls]{cursor:pointer}[tabindex],a,area,button,input,label,select,textarea{-ms-touch-action:manipulation;touch-action:manipulation}[hidden][aria-hidden=false]{clip:rect(0,0,0,0);display:inherit;position:absolute}[hidden][aria-hidden=false]:focus{clip:auto}
</style>
{{ end }}`

const TEMPLATE_STYLE_CANVAS = `{{ define "style_canvas" }}
<style>
body{padding:15px;color:#333;word-break:break-all}
.wrap{border-top:2px solid #666;border-left:2px solid #666}
.row{display:flex;align-items:stretch}
.row:first-child > div{width:20%}
.row:last-child > div{width:50%}
.col{display:flex;flex-direction:column}
.col > div{flex: 1 0 100px}
.content{padding:10px;border-right:2px solid #666;border-bottom:2px solid #666}
</style>
{{ end }}`

const TEMPLATE_STYLE_MARKDOWN = `{{ define "style_markdown" }}
<style>
body{font-family:"Avenir Next",Verdana,"Hiragino Kaku Gothic Pro","Yu Gothic",Meiryo,Osaka,sans-serif;font-size:14px}
h1{font-size:16px;margin-top:0;margin-bottom:10px;word-break:normal}
h2,h3,h4,h5,h6{font-size:14px;margin-top:0;margin-bottom:10px;word-break:normal}
ol,ul{line-height:1.75;margin-top:0;margin-bottom:10px;padding-left:20px}
ol > li{list-style-type:decimal}
li > ol,li > ul{margin-bottom:0;margin-top:0}
li:not(:last-child){margin-bottom:5px}
dl{line-height:1.75;margin-bottom:10px}
dt,dd:not(:last-child){margin-bottom:5px}
hr{margin:10px 0;border:0;height:2px;background-color:#ccc}
table{display:block;margin:10px 0;border-left:solid 1px #ddd}
td,th{border-bottom:solid 1px #ccc;border-right:solid 1px #ccc;padding:8px 10px}
thead td,th{font-weight:700;background-color:#fff;border-top:solid 1px #ddd}
tbody tr:nth-child(odd){background-color:rgba(0,0,0,0.03)}
blockquote{color:#666;margin-left:10px;border-left:solid 4px #ddd;padding-left:10px;margin-top:10px;margin-bottom:10px}
img{vertical-align:middle;margin:10px 0;max-width:100%}
code{border-radius:3px;background-color:#f7f7f7;line-height:1.5;color:#555;font-size:.95em}
pre{border-radius:3px;border:0;margin:10px 0;line-height:1.5;background-color:#f7f7f7;overflow-x:auto;word-wrap:normal;white-space:pre;font-size:13px;font-family:Consolas,"Liberation Mono",Menlo,Courier,monospace;padding:10px}
</style>
{{ end }}`
