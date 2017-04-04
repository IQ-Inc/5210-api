FROM golang:1.8-onbuild

# From the directory of this file, execute
#		docker build -t letsmove .
#
# Docker will construct a go image with this
# source code. Docker builds the source code during
# Image creation. From there, you may start instances
# of the image:
#		docker run -it -p [YOUR_FAVORITE_PORT]:9000 letsmove