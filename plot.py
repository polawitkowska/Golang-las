import pandas as pd
import matplotlib.pyplot as plt
import seaborn as sns

df = pd.read_csv('result.csv')
plt.figure(figsize=(12, 8))

sns.scatterplot(data=df, x='x', y='burnt', hue='p', size='p',
                sizes=(100, 400), alpha=0.6)

plt.title('Forest Fire Simulation Results', fontsize=14)
plt.xlabel('Grid Size', fontsize=12)
plt.ylabel('Percentage of Burnt Area', fontsize=12)
plt.legend(title='Area of trees (%)', title_fontsize=10)

plt.grid(True, linestyle='--', alpha=0.7)

means = df.groupby('x')['burnt'].mean()
plt.plot(means.index, means.values, 'r--', label='Mean burnt area')

plt.tight_layout()
plt.show()