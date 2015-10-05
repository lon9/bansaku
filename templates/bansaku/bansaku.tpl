{% extends "../base.tpl" %}
{% block title %}万策つきたボタン{% endblock %}

{% block content %}
<div data-bind="with: bansaku">
万策尽きた〜:<span data-bind="text: count"></span>
</div>
<button type="button" data-bind="click: send.bind(this)">万策尽きた〜</button>
{% endblock content %}

{% block extrafooter %}
<script src="https://ajax.googleapis.com/ajax/libs/jquery/2.1.4/jquery.min.js"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/jquery-json/2.5.1/jquery.json.min.js"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/knockout/3.3.0/knockout-min.js"></script>
<script src="js/ion.sound.min.js"></script>
<script data-main="js/BansakuMain" src="https://cdnjs.cloudflare.com/ajax/libs/require.js/2.1.20/require.min.js"></script>
{% endblock %}
