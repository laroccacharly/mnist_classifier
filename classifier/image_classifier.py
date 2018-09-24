import tensorflow as tf
from tensorflow import keras
import numpy as np


class ImageClassifier(object):
    # Inspired by https://www.tensorflow.org/tutorials/keras/basic_classification

    def __init__(self):
        # Load data
        self.mnist = tf.keras.datasets.mnist
        (train_images, train_labels), (test_images, test_labels) = self.mnist.load_data()
        # Normalize data from 0 to 255 => 0 to 1
        self.train_images = train_images/255.0
        self.test_images = test_images/255.0
        self.train_labels = train_labels
        self.test_labels = test_labels


        # Basic Dense NN
        self.model = keras.Sequential([
            keras.layers.Flatten(input_shape=(28, 28)),
            keras.layers.Dense(128, activation=tf.nn.relu),
            keras.layers.Dense(10, activation=tf.nn.softmax)
        ])

        self.model.compile(optimizer=keras.optimizers.Adam(),
              loss='sparse_categorical_crossentropy',
              metrics=['accuracy'])

    def evaluate_accuracy_on_test_data(self):
        # returns accuracy (between 0 and 1)
        test_loss, test_acc = self.model.evaluate(self.test_images, self.test_labels)
        return test_acc

    def train(self, epochs):
        self.model.fit(self.train_images, self.train_labels, epochs=epochs)

    def save_model(self, name):
        keras.models.save_model(self.model, name)

    def load_model(self, name):
        self.model = keras.models.load_model(name)

    def classify(self, encoded_image):
        return 7