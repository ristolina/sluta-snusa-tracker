# Generated by Django 4.1.4 on 2022-12-27 15:10

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('slutasnusa', '0002_remove_slutasnusaitem_content'),
    ]

    operations = [
        migrations.AddField(
            model_name='slutasnusaitem',
            name='quitDate',
            field=models.TextField(default='2022-01-01'),
            preserve_default=False,
        ),
    ]
