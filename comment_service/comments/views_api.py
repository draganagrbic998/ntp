from rest_framework import generics, status
from rest_framework.decorators import api_view
from rest_framework.response import Response
from rest_framework_jwt.utils import jwt_decode_handler
from comments.models import Comment, Like
from comments.serializers import CommentSerializer


@api_view(['GET', 'POST'])
def manage_comments(request):
    if request.GET:
        return listing(request)
    return create(request)


def listing(request):
    user_id = jwt_decode_handler(request.headers["Authorization"][4:])["user_id"]
    product_id = request.GET["product"]

    page = 0
    size = 10
    try:
        page = int(request.GET["page"])
    except:
        pass
    try:
        size = int(request.GET["page"])
    except:
        pass

    comments = Comment.objects.filter(product_id=product_id, parent_id=None).order_by('-created_on')[page * size:(page + 1) * size]
    data = []

    for comment in comments:
        data.append({
            'id': comment.id,
            'created_on': comment.created_on,
            'user_id': comment.user_id,
            'email': comment.email,
            'product_id': comment.product_id,
            'text': comment.text,
            'replies': len(Comment.objects.filter(parent_id=comment.id)),
            'likes': len(Like.objects.filter(comment_id=comment.id, dislike=False)),
            'dislikes': len(Like.objects.filter(comment_id=comment.id, dislike=True)),
            'liked': len(Like.objects.filter(user_id=user_id, comment_id=comment.id, dislike=False)) > 0,
            'disliked': len(Like.objects.filter(user_id=user_id, comment_id=comment.id, dislike=True)) > 0
        })

    return Response(data=data, status=status.HTTP_200_OK, content_type='application/json', headers={
        'Access-Control-Expose-Headers': 'First-Page, Last-Page',
        'First-Page': page == 0,
        'Last-Page': (page + 1) * size >= Comment.objects.count()
    })


def create(request):
    comment = Comment.objects.create(product_id=int(request.data['product_id']), text=request.data['text'], parent_id=int(request.data['parent_id']))
    #datum komentara
    comment.user_id = jwt_decode_handler(request.headers["Authorization"][4:])["user_id"]
    comment.email = jwt_decode_handler(request.headers["Authorization"][4:])["email"]
    comment.save()
    return Response(status=status.HTTP_200_OK)


@api_view(['GET'])
def like(request, key):
    user_id = jwt_decode_handler(request.headers["Authorization"][4:])["user_id"]
    disliked = request.GET['dislike'] == "true"
    comment = Comment.objects.get(id=key)

    if not disliked:
        try:
            like = Like.objects.get(user_id=user_id, comment_id=comment.id, dislike=False)
            like.delete()
        except:
            Like.objects.create(user_id=user_id, comment=comment, dislike=False)
            try:
                dislike = Like.objects.get(user_id=user_id, comment_id=comment.id, dislike=True)
                dislike.delete()
            except:
                pass
    else:
        try:
            dislike = Like.objects.get(user_id=user_id, comment_id=comment.id, dislike=True)
            dislike.delete()
        except:
            Like.objects.create(user_id=user_id, comment=comment, dislike=True)
        try:
            like = Like.objects.get(user_id=user_id, comment_id=comment.id, dislike=False)
            like.delete()
        except:
            pass

    return Response(status=status.HTTP_200_OK)


@api_view(['GET'])
def replies(request, key):
    user_id = jwt_decode_handler(request.headers["Authorization"][4:])["user_id"]
    comments = Comment.objects.filter(parent_id=key).order_by('-created_on')
    data = []

    for comment in comments:
        data.append({
            'id': comment.id,
            'created_on': comment.created_on,
            'user_id': comment.user_id,
            'email': comment.email,
            'product_id': comment.product_id,
            'text': comment.text,
            'replies': len(Comment.objects.filter(parent_id=comment.id)),
            'likes': len(Like.objects.filter(comment_id=comment.id, dislike=False)),
            'dislikes': len(Like.objects.filter(comment_id=comment.id, dislike=True)),
            'liked': len(Like.objects.filter(user_id=user_id, comment_id=comment.id, dislike=False)) > 0,
            'disliked': len(Like.objects.filter(user_id=user_id, comment_id=comment.id, dislike=True)) > 0
        })

    return Response(status=status.HTTP_200_OK, data=data, content_type='application/json')


class DeleteComment(generics.DestroyAPIView):
    serializer_class = CommentSerializer
    queryset = Comment.objects.all()

    #ne mogu da prisem ako nije moj komentar...

