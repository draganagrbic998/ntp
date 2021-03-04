from django.test import TestCase, Client


class ActivationTest(TestCase):

    def test_activation(self):
        c = Client()
        response = c.get('/activate/asd')
        self.assertEqual(response.status_code, 200)
        self.assertTemplateUsed(response, 'users/activation.html')