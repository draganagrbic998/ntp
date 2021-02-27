import { Image } from './image';

export interface Advertisement{
    ID: number;
    CreatedOn: string;
    UserId: number;
    Email: string;
    Name: string;
    Category: string;
    Price: number;
    Description: string;
    Images: Image[];
}
