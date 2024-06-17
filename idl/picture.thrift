namespace go image

struct Image {
    1: i64 id
    2: string url
}


struct InsertRequest {
    //1: binary data
}

struct InsertResponse {
    1: Image image
}

struct SearchByImageRequest {
    //1:bindary data
}

struct SearchByTextRequest {
    1: string Text
}

struct SearchResponse {
    1: list<Image> images;
}

service PictureService {
    InsertResponse Insert(1: InsertRequest req) (api.post="picture/insert");
    SearchResponse SearchByText(1: SearchByTextRequest req) (api.get="picture/search/text");
    SearchResponse SearchByImage(1: SearchByImageRequest req) (api.get="picture/search/image");
}