import jwt
from django.test import TestCase, Client

from comment_service import settings
from comments.models import Comment, Like


class CommentsTest(TestCase):

    def setUp(self) -> None:
        self.client = Client()
        self.token = jwt.encode({'user_id': 1, 'email': 'draganaasd@gmail.com'}, settings.SECRET_KEY).decode('utf-8')

        self.comment1 = Comment.objects.create(user_id=1, product_id=1, email='draganaasd@gmail.com', text='comment 1', parent_id=None)
        self.comment2 = Comment.objects.create(user_id=1, product_id=1, email='draganaasd@gmail.com', text='comment 2', parent_id=None)
        self.comment3 = Comment.objects.create(user_id=2, product_id=2, email='dragana.grbic.98@uns.ac.rs', text='comment 3', parent_id=self.comment1.id)

        self.like1 = Like.objects.create(comment_id=self.comment1.id, user_id=1, dislike=False)
        self.like2 = Like.objects.create(comment_id=self.comment1.id, user_id=2, dislike=False)
        self.like3 = Like.objects.create(comment_id=self.comment1.id, user_id=3, dislike=False)

        self.dislike1 = Like.objects.create(comment_id=self.comment1.id, user_id=2, dislike=True)
        self.dislike2 = Like.objects.create(comment_id=self.comment3.id, user_id=1, dislike=True)

    def test_get_comments(self):
        response = self.client.get('/api/comments?product=1', HTTP_AUTHORIZATION=f'JWT {self.token}')
        self.assertEqual(response.status_code, 200)
        data = response.data
        self.assertEqual(len(data), 2)
        self.assertTrue('comment 1' in map(lambda x: x['text'], data))
        self.assertTrue('comment 2' in map(lambda x: x['text'], data))


    def test_get_comments_with_limit(self):
        response = self.client.get('/api/comments?product=1&size=1', HTTP_AUTHORIZATION=f'JWT {self.token}')
        self.assertEqual(response.status_code, 200)
        data = response.data
        self.assertEqual(len(data), 1)

    def test_get_comments_with_offset(self):
        response = self.client.get('/api/comments?product=1&size=1&page=1', HTTP_AUTHORIZATION=f'JWT {self.token}')
        self.assertEqual(response.status_code, 200)
        data = response.data
        self.assertEqual(len(data), 1)

    def test_get_comments_empty(self):
        response = self.client.get('/api/comments?product=1&page=1', HTTP_AUTHORIZATION=f'JWT {self.token}')
        self.assertEqual(response.status_code, 200)
        data = response.data
        self.assertEqual(len(data), 0)

    def test_create_without_parent(self):
        comment = {
            'product_id': 3,
            'text': 'test comment'
        }
        response = self.client.post('/api/comments', HTTP_AUTHORIZATION=f'JWT {self.token}', data=comment, content_type='application/json')

        self.assertEqual(response.status_code, 200)
        self.assertEqual(len(Comment.objects.all()), 4)
        comment = Comment.objects.get(text='test comment')
        self.assertEqual(comment.user_id, 1)
        self.assertEqual(comment.product_id, 3)
        self.assertEqual(comment.email, 'draganaasd@gmail.com')
        self.assertEqual(comment.text, 'test comment')
        self.assertIsNone(comment.parent_id)

    def test_create_with_parent(self):
        comment = {
            'product_id': 3,
            'text': 'test comment',
            'parent_id': 1
        }
        response = self.client.post('/api/comments', HTTP_AUTHORIZATION=f'JWT {self.token}', data=comment, content_type='application/json')

        self.assertEqual(response.status_code, 200)
        self.assertEqual(len(Comment.objects.all()), 4)
        comment = Comment.objects.get(text='test comment')
        self.assertEqual(comment.user_id, 1)
        self.assertEqual(comment.product_id, 3)
        self.assertEqual(comment.email, 'draganaasd@gmail.com')
        self.assertEqual(comment.text, 'test comment')
        self.assertEqual(comment.parent_id, 1)

    def test_delete(self):
        id = Comment.objects.get(text='comment 1').id
        response = self.client.delete(f'/api/comments/{id}', HTTP_AUTHORIZATION=f'JWT {self.token}')
        self.assertEqual(200, response.status_code)
        self.assertEqual(len(Comment.objects.all()), 1)

    def test_delete_not_found(self):
        response = self.client.delete('/api/comments/0', HTTP_AUTHORIZATION=f'JWT {self.token}')
        self.assertEqual(404, response.status_code)
        self.assertEqual(len(Comment.objects.all()), 3)

    def test_delete_forbidden(self):
        id = Comment.objects.get(user_id=2).id
        response = self.client.delete(f'/api/comments/{id}', HTTP_AUTHORIZATION=f'JWT {self.token}')
        self.assertEqual(403, response.status_code)
        self.assertEqual(len(Comment.objects.all()), 3)

    def test_replies(self):
        id = Comment.objects.get(text='comment 1').id
        response = self.client.get(f'/api/comments/{id}/replies', HTTP_AUTHORIZATION=f'JWT {self.token}')
        self.assertEqual(200, response.status_code)
        self.assertEqual(len(response.data), 1)
        self.assertEqual(response.data[0]['text'], 'comment 3')

    def test_replies_empty(self):
        id = Comment.objects.get(text='comment 3').id
        response = self.client.get(f'/api/comments/{id}/replies', HTTP_AUTHORIZATION=f'JWT {self.token}')
        self.assertEqual(200, response.status_code)
        self.assertEqual(len(response.data), 0)

    def test_like(self):
        id = Comment.objects.get(text='comment 3').id
        response = self.client.get(f'/api/comments/{id}/like', HTTP_AUTHORIZATION=f'JWT {self.token}')

        self.assertEqual(200, response.status_code)
        self.assertEqual(len(Like.objects.filter(dislike=False)), 4)
        self.assertEqual(len(Like.objects.filter(dislike=True)), 1)

    def test_unlike(self):
        id = Comment.objects.get(text='comment 1').id
        response = self.client.get(f'/api/comments/{id}/like', HTTP_AUTHORIZATION=f'JWT {self.token}')

        self.assertEqual(200, response.status_code)
        self.assertEqual(len(Like.objects.filter(dislike=False)), 2)
        self.assertEqual(len(Like.objects.filter(dislike=True)), 2)

    def test_dislike(self):
        id = Comment.objects.get(text='comment 1').id
        response = self.client.get(f'/api/comments/{id}/like?dislike=true', HTTP_AUTHORIZATION=f'JWT {self.token}')

        self.assertEqual(200, response.status_code)
        self.assertEqual(len(Like.objects.filter(dislike=False)), 2)
        self.assertEqual(len(Like.objects.filter(dislike=True)), 3)

    def test_undislike(self):
        id = Comment.objects.get(text='comment 3').id
        response = self.client.get(f'/api/comments/{id}/like?dislike=true', HTTP_AUTHORIZATION=f'JWT {self.token}')

        self.assertEqual(200, response.status_code)
        self.assertEqual(len(Like.objects.filter(dislike=False)), 3)
        self.assertEqual(len(Like.objects.filter(dislike=True)), 1)

    def test_statistic_comments_invalid(self):
        response = self.client.get(f'/api/statistic-comments/2020/2020')
        self.assertEqual(response.status_code, 400)

    def test_statistic_comments(self):
        response = self.client.get(f'/api/statistic-comments/2020/2021')
        self.assertEqual(response.status_code, 200)
        self.assertEqual(response.data, [[2020, 0], [2021, 3]])

    def test_statistic_likes_invalid(self):
        response = self.client.get(f'/api/statistic-likes/2020/2020')
        self.assertEqual(response.status_code, 400)

    def test_statistic_likes(self):
        response = self.client.get(f'/api/statistic-likes/2020/2021')
        self.assertEqual(response.status_code, 200)
        self.assertEqual(response.data, [[2020, 0], [2021, 3]])

    def test_statistic_dislikes_invalid(self):
        response = self.client.get(f'/api/statistic-dislikes/2020/2020')
        self.assertEqual(response.status_code, 400)

    def test_statistic_dislikes(self):
        response = self.client.get(f'/api/statistic-dislikes/2020/2021')
        self.assertEqual(response.status_code, 200)
        self.assertEqual(response.data, [[2020, 0], [2021, 2]])
