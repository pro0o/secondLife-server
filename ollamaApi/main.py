from flask import Flask, request, jsonify
import ollama

app = Flask(__name__)

@app.route('/recyclingData/tips', methods=['POST'])
def chat():
    if request.is_json:
        content = request.json.get('content', '')
        response = ollama.chat(model='llama3', messages=[{'role': 'user', 'content': content}])

        response_content = response['message']['content']

        return jsonify({'response': response_content})
    else:
        return jsonify({'error': 'Request must be in JSON format'}), 400

if __name__ == '__main__':
    app.run(host='localhost', debug=True)
