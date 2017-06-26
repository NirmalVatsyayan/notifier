import { NotificationProjPage } from './app.po';

describe('notification-proj App', () => {
  let page: NotificationProjPage;

  beforeEach(() => {
    page = new NotificationProjPage();
  });

  it('should display message saying app works', () => {
    page.navigateTo();
    expect(page.getParagraphText()).toEqual('app works!');
  });
});
