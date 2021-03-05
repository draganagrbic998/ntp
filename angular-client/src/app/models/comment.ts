export interface Comment{
    id: number;
    created_on: string;
    user_id: number;
    product_id: number;
    email: string;
    text: string;
    parent_id: number;
    likes: number;
    dislikes: number;
    liked: boolean;
    disliked: boolean;
}
