# app.py
from datetime import datetime
from flask import Flask, request
from flask_sqlalchemy import SQLAlchemy
from migrate.versioning.api import upgrade, version_control

app = Flask(__name__)

url = 'sqlite:///beerstore.db'
# if app.config['FLASK_ENV'] == 'development':
#     url ="postgres://postgres:postgres@127.0.0.1/beerstore"
app.config['SQLALCHEMY_DATABASE_URI'] = url
app.config['SQLALCHEMY_TRACK_MODIFICATIONS'] = False

db = SQLAlchemy(app)

# if app.config['FLASK_ENV'] == 'development':
try:
    version_control(url, 'migrations')
except:
    print("db already versioned!")

upgrade(repository='migrations', url=url, debug='False')


class Media(db.Model):
    __tablename__ = "media"
    idmedia = db.Column(db.Integer, nullable=False, primary_key=True)
    creationdatemedia = db.Column(
        db.DateTime, nullable=False, default=datetime.now)
    datamedia = db.Column(db.LargeBinary, nullable=False)
    nomemedia = db.Column(db.String(255), nullable=False)
    mimemedia = db.Column(db.String(255), nullable=False)


class Beer(db.Model):
    __tablename__ = "beer"
    idbeer = db.Column(db.Integer, primary_key=True)
    creationdatebeer = db.Column(
        db.DateTime, nullable=False, default=datetime.now)
    titlebeer = db.Column(db.String(255), nullable=False)
    descriptionbeer = db.Column(db.Text)
    idmedia = db.Column(db.Integer, db.ForeignKey('media.idmedia'))


@app.route("/status")
def hello():
    return "ONLINE!"


@app.route("/beer/list", methods=["GET"])
def beerlist():
    page = request.args.get("page")
    pageSize = request.args.get("pageSize")
    return Beer.query.paginate(page, pageSize).items
