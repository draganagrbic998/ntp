import { Image } from './image';

export interface Advertisement{
    ID: number;
    CreatedOn: string;
    UserID: number;
    Name: string;
    Category: string;
    Price: string;
    Description: string;
    Images: Image[];
}
