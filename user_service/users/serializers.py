from rest_framework import serializers

from users.models import CustomUser


class RegistrationSerializer(serializers.Serializer):
    email = serializers.EmailField(max_length=100)
    password = serializers.CharField(max_length=100)
    first_name = serializers.CharField(max_length=100)
    last_name = serializers.CharField(max_length=100)
    address = serializers.CharField(max_length=100)
    city = serializers.CharField(max_length=100)
    zip_code = serializers.CharField(max_length=100)

    def create(self, validated_data):
        return CustomUser()

    def update(self, instance, validated_data):
        return instance


class ProfileSerializer(serializers.Serializer):
    email = serializers.EmailField(max_length=100)
    first_name = serializers.CharField(max_length=100)
    last_name = serializers.CharField(max_length=100)
    address = serializers.CharField(max_length=100)
    city = serializers.CharField(max_length=100)
    zip_code = serializers.CharField(max_length=100)
    old_password = serializers.CharField(allow_null=True, allow_blank=True, max_length=100)
    new_password = serializers.CharField(allow_null=True, allow_blank=True, max_length=100)

    def create(self, validated_data):
        return CustomUser()

    def update(self, instance, validated_data):
        return instance
