#!/bin/sh
set -e

# Workaround old docker images with incorrect $HOME
# check https://github.com/docker/docker/issues/2968 for details
if [ "${HOME}" = "/" ]
then
  export HOME=$(getent passwd $(id -un) | cut -d: -f6)
fi

replace='https://github.com/'
to_replace='git@github.com:'
NEW_CIRCLE_REPOSITORY_URL="${CIRCLE_REPOSITORY_URL/$to_replace/$replace}"

if [ -e /home/user/work/src/github.com/Nordgedanken/Morpheusv2/.git ]
then
  cd /home/user/work/src/github.com/Nordgedanken/Morpheusv2/
  git remote set-url origin "$NEW_CIRCLE_REPOSITORY_URL" || true
else
  mkdir -p /home/user/work/src/github.com/Nordgedanken/Morpheusv2/
  cd /home/user/work/src/github.com/Nordgedanken/Morpheusv2/
  git clone "$NEW_CIRCLE_REPOSITORY_URL" .
fi

if [ -n "$CIRCLE_TAG" ]
then
  git fetch --force origin "refs/tags/${CIRCLE_TAG}"
else
  git fetch --force origin "${CIRCLE_BRANCH}:remotes/origin/${CIRCLE_BRANCH}"
fi


if [ -n "$CIRCLE_TAG" ]
then
  git reset --hard "$CIRCLE_SHA1"
  git checkout -q "$CIRCLE_TAG"
elif [ -n "$CIRCLE_BRANCH" ]
then
  git reset --hard "$CIRCLE_SHA1"
  git checkout -q -B "$CIRCLE_BRANCH"
fi

git reset --hard "$CIRCLE_SHA1"
