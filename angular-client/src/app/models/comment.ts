export interface Comment{
    id: number;
    created_on: string;
    user_id: number;
    email: string;
    product_id: number;
    text: string;
    parent_id: number;
    likes: number;
    dislikes: number;
    liked: boolean;
    disliked: boolean;
}
