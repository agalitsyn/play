#!/usr/bin/python2
# -*- coding: utf-8 -*-

import os
import logging

from slugify import slugify

LOG = logging.getLogger('{}:{}'.format(__file__, __name__))
logging.basicConfig(level=logging.INFO)

ROOT_DIR = '/Users/anton/Google Drive/Catalog'
PRJ_DIR = '/Users/anton/Projects/catalog/src/ui/app/media'

ALLOWED_EXTENSIONS = ('.jpg',)
IGNORED_FOLDERS = ('no copy',
                   'no copy 2')


def main():
    for dirpath, dirnames, filenames in os.walk(ROOT_DIR):
        dirnames[:] = [d for d in dirnames if d not in IGNORED_FOLDERS]
        filenames[:] = [name for name in filenames
                        if name.lower().endswith(ALLOWED_EXTENSIONS)]
        if not filenames:
            continue

        folder_name = slugify(os.path.relpath(dirpath, ROOT_DIR))
        dst = os.path.join(PRJ_DIR, folder_name)
        try:
            os.mkdir(dst, 0755)
        except OSError as e:
            LOG.debug(e)
        else:
            LOG.info('Created {}'.format(dst))

if __name__ == '__main__':
    main()
