import pandas as pd
import matplotlib.pyplot as plt
import numpy as np
from scipy.interpolate import griddata

df = pd.read_csv('results_with_humidity.csv')

df['grid_size'] = df['x'] * df['y']

grid_size_min, grid_size_max = df['grid_size'].min(), df['grid_size'].max()
p_min, p_max = df['p'].min(), df['p'].max()
grid_x, grid_y = np.mgrid[grid_size_min:grid_size_max:100j, p_min:p_max:100j]

points = np.column_stack((df['grid_size'], df['p']))
values = df['burnt']
grid_z = griddata(points, values, (grid_x, grid_y), method='cubic')

fig = plt.figure(figsize=(12, 8))
ax = fig.add_subplot(111, projection='3d')

surface = ax.plot_surface(grid_x, grid_y, grid_z, 
                         cmap='viridis',
                         alpha=0.8,
                         linewidth=0)

ax.set_title('Forest Fire Simulation Results', fontsize=14)
ax.set_xlabel('Grid Size (x*y)', fontsize=12)
ax.set_ylabel('Trees density (%)', fontsize=12)
ax.set_zlabel('Percentage of Burnt Area', fontsize=12)

colorbar = plt.colorbar(surface)
colorbar.set_label('Burnt Area (%)', fontsize=10)

ax.view_init(elev=16, azim=-57)

plt.tight_layout()
plt.show()