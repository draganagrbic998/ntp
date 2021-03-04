from django.test import TestCase, Client


class RegistrationTest(TestCase):

    def setUp(self) -> None:
        self.client = Client()
        self.registration = {
            'email': 'draganaasd@gmail.com',
            'password': 'asd',
            'first_name': 'Dragana',
            'last_name': 'Grbic',
            'address': 'Lasla Gala 23',
            'city': 'Novi Sad',
            'zip_code': '21000'
        }

    def test_registration(self):
        response = self.client.post('/api/register', data=self.registration, content_type='application/json')
        self.assertEqual(response.status_code, 200)
        self.assertTrue('activation_link' in response.data)
        return response

    def test_bad_email(self):
        registration = {
            'email': 'draganaasd',
            'password': 'asd',
            'first_name': 'Dragana',
            'last_name': 'Grbic',
            'address': 'Lasla Gala 23',
            'city': 'Novi Sad',
            'zip_code': '21000'
        }
        response = self.client.post('/api/register', data=registration, content_type='application/json')
        self.assertEqual(response.status_code, 400)

    def test_empty_fields(self):
        registration = {
            'email': '',
            'password': '',
            'first_name': '',
            'last_name': '',
            'address': '',
            'city': '',
            'zip_code': ''
        }
        response = self.client.post('/api/register', data=registration, content_type='application/json')
        self.assertEqual(response.status_code, 400)

    def test_missing_fields(self):
        registration = {}
        response = self.client.post('/api/register', data=registration, content_type='application/json')
        self.assertEqual(response.status_code, 400)

    def test_login(self):
        response = self.test_registration()

        response = self.client.get(f'/activate/{response.data["activation_link"]}')
        self.assertEqual(response.status_code, 200)
        self.assertTemplateUsed(response, 'users/activation.html')

        status, _, _ = authenticate(self.client, self.registration['email'], self.registration['password'])
        self.assertEqual(status, 200)

    def test_login_wrong_password(self):
        response = self.test_registration()

        response = self.client.get(f'/activate/{response.data["activation_link"]}')
        self.assertEqual(response.status_code, 200)
        self.assertTemplateUsed(response, 'users/activation.html')

        status, user_id, token = authenticate(self.client, self.registration['email'], 'hehe')
        self.assertEqual(status, 400)

    def test_login_without_activation(self):
        self.test_registration()

        status, user_id, token = authenticate(self.client, self.registration['email'], self.registration['password'])
        self.assertEqual(status, 400)


class ProfileTest(TestCase):

    def setUp(self) -> None:
        self.client = Client()
        self.registration = {
            'email': 'draganaasd@gmail.com',
            'password': 'asd',
            'first_name': 'Dragana',
            'last_name': 'Grbic',
            'address': 'Lasla Gala 23',
            'city': 'Novi Sad',
            'zip_code': '21000'
        }

        response = self.client.post('/api/register', data=self.registration, content_type='application/json')
        self.client.get(f'/activate/{response.data["activation_link"]}')
        _, _, token = authenticate(self.client, self.registration['email'], self.registration['password'])
        self.token = token


    def test_profile_update(self):
        profile = {
            'email': 'dragana.grbic.98@uns.ac.rs',
            'first_name': 'Imence',
            'last_name': 'Prezimence',
            'address': 'Adresica',
            'city': 'Gradic',
            'zip_code': 'zip_code',
            'old_password': 'asd',
            'new_password': 'qwe'
        }
        response = self.client.post('/api/update-profile', data=profile, HTTP_AUTHORIZATION=f'JWT {self.token}')

        self.assertEqual(response.status_code, 200)
        self.assertEqual(profile['email'], response.data['email'])
        self.assertEqual(profile['first_name'], response.data['first_name'])
        self.assertEqual(profile['last_name'], response.data['last_name'])
        self.assertEqual(profile['address'], response.data['address'])
        self.assertEqual(profile['city'], response.data['city'])
        self.assertEqual(profile['zip_code'], response.data['zip_code'])

        _, _, token = authenticate(self.client, profile['email'], profile['new_password'])

        new_profile = dict()
        new_profile.update(self.registration)
        new_profile.update({
            'old_password': profile['new_password'],
            'new_password': self.registration['password']
        })
        response = self.client.post('/api/update-profile', data=new_profile, HTTP_AUTHORIZATION=f'JWT {token}')
        self.assertEqual(response.status_code, 200)
        self.token = response.data['token']


    def test_bad_email(self):
        profile = {
            'email': 'dragana.grbic.98',
            'first_name': 'Imence',
            'last_name': 'Prezimence',
            'address': 'Adresica',
            'city': 'Gradic',
            'zip_code': 'zip_code',
            'old_password': 'asd',
            'new_password': 'qwe'
        }
        response = self.client.post('/api/update-profile', data=profile, HTTP_AUTHORIZATION=f'JWT {self.token}')
        self.assertEqual(response.status_code, 400)


    def test_empty_fields(self):
        profile = {
            'email': '',
            'first_name': '',
            'last_name': '',
            'address': '',
            'city': '',
            'zip_code': '',
            'old_password': '',
            'new_password': ''
        }
        response = self.client.post('/api/update-profile', data=profile, HTTP_AUTHORIZATION=f'JWT {self.token}')
        self.assertEqual(response.status_code, 400)

    def test_missing_fields(self):
        profile = {}
        response = self.client.post('/api/update-profile', data=profile, HTTP_AUTHORIZATION=f'JWT {self.token}')
        self.assertEqual(response.status_code, 400)

    def test_unauthorized(self):
        profile = {}
        response = self.client.post('/api/update-profile', data=profile)
        self.assertEqual(response.status_code, 401)

    def test_wrong_old_password(self):
        profile = {
            'email': 'dragana.grbic.98@uns.ac.rs',
            'first_name': 'Imence',
            'last_name': 'Prezimence',
            'address': 'Adresica',
            'city': 'Gradic',
            'zip_code': 'zip_code',
            'old_password': 'asdd',
            'new_password': 'qwe'
        }
        response = self.client.post('/api/update-profile', data=profile, HTTP_AUTHORIZATION=f'JWT {self.token}')
        self.assertEqual(response.status_code, 400)


def authenticate(client: Client, username, password):
    response = client.post('/api/login', {
        'username': username,
        'password': password
    })
    if response.status_code == 200:
        return response.status_code, response.data['id'], response.data['token']
    return response.status_code, None, None



