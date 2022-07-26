from uuid import uuid4

import requests
from loguru import logger


class TestBlockChain(object):
    req_obj = requests.Session()
    base_url = "http://127.0.0.1:5000"

    def test_mine(self):
        """挖矿"""
        res = self.req_obj.request(method='get', url=f'{self.base_url}/mine')
        logger.info(res.json())
        # 只有一笔交易，即挖矿的奖励给自己的
        assert len(res.json()['transactions']) == 1

    def test_send_transactions(self):
        sender = str(uuid4()).replace('-', '')
        data = {'sender': sender, 'recipient': sender, 'amount': 1001010, }
        res = self.req_obj.request(method='post', url=f'{self.base_url}/transactions/new', json=str(data))
        logger.info(res.json())

        res = self.req_obj.request(method='get', url=f'{self.base_url}/mine')
        logger.info(res.json())

    def test_chain(self):
        res = self.req_obj.request(method='get', url=f'{self.base_url}/chain')
        logger.info(res.json())


if __name__ == '__main__':
    a = TestBlockChain()
    a.test_chain()
