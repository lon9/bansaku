<html>
<head>
<meta charset="utf-8">
<meta name="viewport" name="width=device-width, initial-scale=1.0">
<title>{% block title %}Zepher{% endblock %}</title>
{% block extrahead %}{% endblock %}
</head>
<body>
{% block content %}
  {{ content }}
{% endblock %}
{% block extrafooter %}{% endblock %}
</body>
</html>
