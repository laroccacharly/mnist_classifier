from samples_generator import SamplesGenerator


def test_sample_generation():
    samples = SamplesGenerator.get_samples(10)
    assert(len(samples) == 10)
    print(samples[0])
    assert(isinstance(samples[0]["encodedImage"], str))
    assert(isinstance(samples[0]["label"], int))
