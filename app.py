from flask import Flask, render_template, Response, request, send_file
import pandas as pd
import os


app = Flask(__name__)


@app.route("/")
def index():
    return render_template("index.html")


@app.route("/api/data")
def get_data():
    df = pd.read_csv("./static/data/KARMY-_punkter.csv")
    df['WKT'] = df['WKT'].str[7:-2].str.split(' ')

    return Response(
            df.to_json(orient="records"),
            mimetype="application/json"
        )


@app.route("/api/images")
def get_image():
    image_id = request.args.get("id")
    if not image_id:
        return send_file("./static/images/placeholder.jpg")

    if not os.path.isfile(f"./static/images/{image_id}.jpg"):
        return send_file("./static/images/placeholder.jpg")

    return send_file(f"./static/images/{image_id}.jpg", mimetype="image/jpeg")
