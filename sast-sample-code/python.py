import os
import urllib
from flask import Flask, request
from django.db import connection, models
from django.db.models.expressions import RawSQL

app = Flask(__name__)

@app.route("/code-execution")
def code_execution():
    code1 = request.args.get("code1")
    exec("setname('%s')" % code1)
    return a

@app.route("/open-redirect")
def open_redirect():
    redirect_loc = request.args.get('redirect')
    return redirect(redirect_loc)


@app.route("/sqli/<username>")
def show_user(username):
    with connection.cursor() as cursor:
      cursor.execute("SELECT * FROM users WHERE username = '%s'" % username)

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=9000)
