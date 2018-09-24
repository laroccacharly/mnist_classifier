from image_classifier import ImageClassifier
import pytest
import os

def test_classifier():
    # Payload from request 
    payload = {"encoded_image": "image123"}
    # Classify the image, returns a label (int)
    classifier = ImageClassifier()
    label = classifier.classify(payload["encoded_image"])
    assert(isinstance(label, int))

def test_train():
    # evaluate the classifier with test data
    classifier = ImageClassifier()
    accuracy_no_training = classifier.evaluate_accuracy_on_test_data()
    # should return bad accuracy because not trained yet 
    assert(accuracy_no_training < 0.3)

    # train classifier 
    classifier.train(epochs=1)
    # evaluate the classifier again 
    accuracy_with_some_training = classifier.evaluate_accuracy_on_test_data()
    # check if accuracy is better 
    assert(accuracy_with_some_training > accuracy_no_training)
    # Save classifier model in a file
    file_name = "test_model"
    classifier.save_model(name=file_name)
    # evaluate classifier with loaded file 
    new_classifier = ImageClassifier()
    new_classifier.load_model(name=file_name)
    # check if accuracy the same as before 
    accuracy_of_loaded_model = new_classifier.evaluate_accuracy_on_test_data()
    assert pytest.approx(accuracy_of_loaded_model, 0.2) == accuracy_with_some_training
    # clean up file
    os.remove(file_name)



