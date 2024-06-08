from flask import Flask, render_template, Response
import pandas as pd


app = Flask(__name__)


@app.route("/")
def index():
    return render_template("index.html")


@app.route("/api/data")
def data():
    df = pd.read_csv("./static/data/KARMY-_punkter.csv")
    df['WKT'] = df['WKT'].str[7:-2].str.split(' ')

    return Response(
            df.to_json(orient="records"),
            mimetype="application/json"
        )
