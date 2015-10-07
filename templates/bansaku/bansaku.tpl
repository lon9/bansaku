{% extends "../base.tpl" %}
{% block title %}万策つきたボタン{% endblock %}
{% block extrahead %}
<link href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/css/bootstrap.min.css" rel="stylesheet">
<link href="css/custom.css" rel="stylesheet">
{% endblock %}

{% block content %}
<div class="outer">
<div class="container">
<div class="row">
<div class="col-xs-12 col-sm-offset-2 col-sm-8 centerd">
<div data-bind="with: bansaku">
<h1 data-bind="text: count"></h1>
</div>
<button type="button" class="btn btn-primary" data-bind="click: send.bind(this)">万策尽きた〜</button>
</div>
</div>
</div>
</div>
{% endblock content %}

{% block extrafooter %}
<script src="https://ajax.googleapis.com/ajax/libs/jquery/2.1.4/jquery.min.js"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/jquery-json/2.5.1/jquery.json.min.js"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/knockout/3.3.0/knockout-min.js"></script>
<script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/js/bootstrap.min.js"></script>
<script src="js/ion.sound.min.js"></script>
<script data-main="js/BansakuMain" src="https://cdnjs.cloudflare.com/ajax/libs/require.js/2.1.20/require.min.js"></script>
{% endblock %}
