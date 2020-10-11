import { TestBed } from '@angular/core/testing';

import { MuchoAnswerService } from './mucho-answer.service';

describe('MuchoAnswerService', () => {
  let service: MuchoAnswerService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(MuchoAnswerService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
