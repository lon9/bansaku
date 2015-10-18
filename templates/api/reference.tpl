{% extends "../base.tpl" %}
{% block title %}万策つきたボタン リファレンス{% endblock %}
{% block extrahead %}
<link href="https://cdnjs.cloudflare.com/ajax/libs/skeleton/2.0.4/skeleton.min.css" rel="stylesheet">
<link href="css/custom.css" rel="stylesheet">
{% endblock %}

{% block content %}
<div class="container">
  <div class="row">
    <div class="nine columns">
      <h3>万策尽きたボタンAPIリファレンス</h3>
      <p>万策ボタンのAPIです。ご自由にお使いください。</p>
      <h4>レートリミット</h4>
      <p>10 request / 1 second</p>
      <div id="get">
        <h5>万策回数取得</h5>
        <h6>URL</h6>
        <pre><code>http://bansaku.zepher.me/api/count</code></pre>
        <h6>レスポンス</h6>
        <pre><code>{
  "count": 100
}</code></pre>
      </div>
    </div>
    <div class="three columns">
      <div class="container">
        <ul>
          <li><a href="#get">万策回数取得</a></li>
        </ul>
      </div>
    </div>
</div>
{% endblock content %}
