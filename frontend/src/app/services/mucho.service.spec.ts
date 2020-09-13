import { TestBed } from '@angular/core/testing';

import { MuchoService } from './mucho.service';

describe('MuchoService', () => {
  let service: MuchoService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(MuchoService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
