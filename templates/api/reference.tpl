{% extends "../base.tpl" %}
{% block title %}万策つきたボタン{% endblock %}
{% block extrahead %}
<link href="https://cdnjs.cloudflare.com/ajax/libs/skeleton/2.0.4/skeleton.min.css" rel="stylesheet">
{% endblock %}

{% block content %}
<div class="container">
  <h4>万策尽きたボタンAPIリファレンス</h4>
  <p>万策ボタンのAPIです。ご自由にお使いください。</p>
  <h5>レートリミット</h5>
  <p>10 request / 1 second</p>
  <h5>万策回数取得</h5>
  <pre><code>https://bansaku.zepher.me/api/count/</code></pre>
</div>
{% endblock content %}
