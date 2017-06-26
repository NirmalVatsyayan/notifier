from mongoengine import *
import datetime

connect('notify')

class Notification(DynamicDocument):
    header = StringField(required=True, max_length=150)
    payload = StringField(required=True, max_length=300)
    imageurl = StringField(required=True, max_length=150)
    userquery = StringField(required=True, max_length=150)
    notification_time = DateTimeField(default=datetime.datetime.utcnow)



'''
Celery integration
'''
def send_user_notification(user_id, notification_payload):
    print(notification_payload)
    return True



'''
TO-DO Get more elegant way to extract user ids
'''
def send_notification(notification_data):
    try: 
        query = notification_data.userquery
        query_list = query.split('(')
        query_list = query_list[1].split(')')
        user_id = query_list[0].split(',')

        payload = {}
        payload["header"] = notification_data.header
        payload["content"] = notification_data.payload
        payload["image_url"] = notification_data.imageurl

        for user in user_id:
            send_user_notification(user_id, payload)
    except:
        pass


'''
There must be a driver function to filter from Notification
collection depending upon the time of notification
''' 
for data in Notification.objects.all():
    send_notification(data)

