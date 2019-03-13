# app.py
from os import environ 
from flask_cors import CORS
from datetime import datetime
from flask import Flask, request, jsonify
from flask_sqlalchemy import SQLAlchemy
from migrate.versioning.api import upgrade, version_control

app = Flask(__name__)
CORS(app)

url = 'sqlite:///beerstore.db'
if 'DATABASE_URL' in environ:
    url = environ['DATABASE_URL']
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

    def to_dict(self):
        return {
            "idbeer": self.idbeer,
            "creationdatebeer": self.creationdatebeer,
            "titlebeer": self.titlebeer,
            "descriptionbeer": self.descriptionbeer,
            "idmedia": self.idmedia,
        }


@app.route("/status")
def hello():
    return "ONLINE!"


@app.route("/beer/list", methods=["GET"])
def beerlist():
    page = request.args.get("page", default=1, type=int)
    pageSize = request.args.get("pageSize", default=10, type=int)
    search = request.args.get("search", default="")
    items = Beer.query.filter(Beer.titlebeer.contains(
        search)).paginate(page, pageSize).items
    return jsonify([item.to_dict() for item in items])

@app.route("/beer/<idbeer>")
def beerfind(idbeer):
    beer = Beer.query.get(idbeer)
    return jsonify(beer.to_dict())

