from flask import Flask, request, jsonify
from ultralytics import YOLO
import cv2
import numpy as np
import os
import requests
from io import BytesIO

app = Flask(__name__)
model = YOLO("yolov8n.pt")

@app.route('/detect', methods=['POST'])
def detect_objects():
    data = request.get_json()
    if 'image_url' not in data:
        return jsonify({"error": "No image URL provided"}), 400
    
    image_url = data['image_url']
    response = requests.get(image_url)
    if response.status_code != 200:
        return jsonify({"error": "Failed to fetch image from URL"}), 400
    
    image_arr = np.frombuffer(BytesIO(response.content).read(), np.uint8)
    img = cv2.imdecode(image_arr, cv2.IMREAD_COLOR)
    
    results = model(img)
    
    detections = []
    for r in results:
        for box in r.boxes:
            x1, y1, x2, y2 = map(int, box.xyxy[0])
            conf = float(box.conf[0])
            cls = int(box.cls[0])
            detections.append({
                "class": model.names[cls],
                "confidence": conf,
                "bbox": [x1, y1, x2, y2]
            })
    
    return jsonify({"detections": detections})

if __name__ == '__main__':
    app.run(debug=True)
