from flask import Flask
from flask import jsonify, request, abort
from samples_generator import SamplesGenerator
from image_classifier import ImageClassifier

app = Flask(__name__)


@app.route("/samples")
def handle_samples():
    payload = SamplesGenerator.get_samples(10)
    return jsonify(payload)

@app.route("/classifySample", methods=['POST'])
def handle_classify():
    app.logger.info(request.get_data())
    app.logger.info(request.headers)
    app.logger.info(request.json)
    classifier = ImageClassifier()
    label = classifier.classify(request.json['encoded_image'])
    return jsonify({"label": label})
    

if __name__ == "__main__":
    app.run(host='0.0.0.0', port=80, debug=True)