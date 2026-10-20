## Learner APIs

### Get data to populate users dashboard
> GET /v1/learner/dashboard





### Will fetch all of the learning paths available
> GET /v1/learner/all-learning-paths

Response:
```
{
    [
        {
            id: 1,
            title: "Introduction to ML",
            description: "Test description",
            thumbnailPath: "www.thumbnail.com"
            trackOwner: "Dr. Sarah Chen",
            videoIds: ['123', '456', '789'],
            quizIds: ['1234', '5643', '2124']
            requiredQuizIds: ['2124'],
            contentOrder: ['123', '1234', '456', '5643', '789', '2124']
            estimatedDurationTocomplete: "8 hours",
            totalEnrolled: 536,
            avgRating: 4.7,
            tags: ["AI", "Machine Learning", "Python"]
        },
        {
            id: 2,
            title: "Advanced React Patterns",
            description: "Test description",
            thumbnailPath: "www.thumbnail.com"
            trackOwner: "Dr. Sarah Chen",
            videoIds: ['123', '456', '789'],
            quizIds: ['1234', '5643', '2124']
            requiredQuizIds: ['2124'],
            contentOrder: ['123', '1234', '456', '5643', '789', '2124']
            estimatedDurationTocomplete: "8 hours",
            totalEnrolled: 536,
            avgRating: 4.7,
            tags: ["React", "Javascript", "Frontend"]
        },
    ]
}
```


### Fetches details the learnings for current logged in user

> Needed for `MyLearnings` page; list down user's current learnings

GET `/v1/learner/my-learnings`

Response:
```
{
    [
        {   
            id: "234"
            learningTrackId: 1,
            
        },
        {   
            learningTrackId: 987,
            title: "React design patterns",
            description "Description of path",
            thumnailURL: "www.link.com",
            videoIds: ['123', '456', '789'],
            quizIds: ['1234', '5643', '2124'],
            trackOwner: "Dr. Sarah Chen",
            totalVideoDuration: "2h 48m"
            learnerTrackStatus: {
                currentTrackStatus: "Completed"
                completedVideoDuration: "2h 48m",
                currentItemId: "789"
            }
        },
    ]
}
```


### Learning Page API

> This will fetch the learning track detail with the video playback capibility

GET: `/v1/learner/video?videoId=123`

Response:

```
{
    id: 123
    title: "Video title"
    description: "Video description"
    duration: 82718 // In secs
    videoRequired: true // [true / false]
    hasResources: true // [true / false]
    files: [

    ],
    videoStatus: {
        id: 1,
        videoId: 123,
        status: "In Progress" // Not started / In Progress / Completed
        lastViewedDuration: 8218 // In secs
        trackStatusId: xyz
    }
}
```


> Stream video

GET: `/v1/learner/video/:videoId/playback/:segment`


### Fetches details about logged in user's learning path
> GET /v1/learner/my-profile

Response:
```
{
    name: "Swaniket Chowdhury",
    email: "test@email.com",
    memberSince: "Oct 2025",
    enrolled: 2,
    completed: 2,
    totalHoursLearned: "14h 30m"
}
```

### Edit user's profile
> PATCH /v1/learner/my-profile

Payload:
```
{
    name: "Updated Name",
}
```

Response:
```
{
    name: "Updated Name",
}
```




