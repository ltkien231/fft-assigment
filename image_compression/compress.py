from matplotlib.image import imread
import numpy as np
import matplotlib.pyplot as plt
import os
from mpl_toolkits.mplot3d import Axes3D

plt.rcParams['figure.figsize'] = [12, 8]
plt.rcParams.update({'font.size': 18})

# Show original image
A = imread(os.path.join('.', 'images', 'child.jpeg'))
Abw = np.mean(A, -1)  # Convert RGB to grayscale so that it's simpler to work

plt.imshow(Abw, cmap='gray')
plt.axis('off')
plt.savefig('./images/origin.jpeg', bbox_inches='tight')
plt.show()

# keep is % of big coefficient we want to keep
for keep in [0.1, 0.01, 0.002]:
    # Compute FFT of image using fft2
    At = np.fft.fft2(Abw)
    F = np.log(np.abs(np.fft.fftshift(At))+1)  # Put FFT on log scale
    plt.imshow(F, cmap='gray')

    # Zero out all small coefficients and inverse transform
    Bt = np.sort(np.abs(np.reshape(At, -1)))

    thresh = Bt[int(np.floor((1-keep)*len(Bt)))]
    ind = np.abs(At) > thresh
    Atlow = At * ind
    Flow = np.log(np.abs(np.fft.fftshift(Atlow))+1)  # Put FFT on log scale

    # Plot reconstruction image
    Alow = np.fft.ifft2(Atlow).astype('uint8')

    plt.imshow(Alow, cmap='gray')
    plt.axis('off')

    plt.savefig('./images/compress_keep_'+str(keep) +
                '.jpeg', bbox_inches='tight')
    plt.show()
